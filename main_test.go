package main

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	router "github.com/v2fly/v2ray-core/v5/app/router/routercommon"
	"google.golang.org/protobuf/proto"
)

func TestParseEntry(t *testing.T) {
	testCases := []struct {
		name      string
		typ       string
		rule      string
		wantPlain string
		wantAffs  []string
		wantErr   bool
	}{
		{name: "domain", typ: "domain", rule: "Example.COM", wantPlain: "domain:example.com"},
		{name: "sorted attrs", typ: "full", rule: "a.example.com @cn @ads", wantPlain: "full:a.example.com:@ads,@cn"},
		{name: "duplicated attrs", typ: "domain", rule: "example.com @ads @ads", wantPlain: "domain:example.com:@ads"},
		{name: "affiliations", typ: "domain", rule: "example.com &other @ads", wantPlain: "domain:example.com:@ads", wantAffs: []string{"other"}},
		{name: "regexp", typ: "regexp", rule: `^example\.com$`, wantPlain: `regexp:^example\.com$`},
		{name: "tld", typ: "domain", rule: "google", wantPlain: "domain:google"},
		{name: "hyphen inside", typ: "domain", rule: "my-example.com", wantPlain: "domain:my-example.com"},
		{name: "keyword allows partial", typ: "keyword", rule: "-ads-", wantPlain: "keyword:-ads-"},
		{name: "invalid regexp", typ: "regexp", rule: "^example(", wantErr: true},
		{name: "empty rule", typ: "domain", rule: "   ", wantErr: true},
		{name: "unknown type", typ: "prefix", rule: "example.com", wantErr: true},
		{name: "invalid domain", typ: "domain", rule: "exa_mple.com @ads", wantErr: true},
		{name: "empty label", typ: "domain", rule: "example..com", wantErr: true},
		{name: "trailing dot", typ: "full", rule: "example.com.", wantErr: true},
		{name: "leading dot", typ: "domain", rule: ".example.com", wantErr: true},
		{name: "leading hyphen", typ: "domain", rule: "-example.com", wantErr: true},
		{name: "trailing hyphen", typ: "full", rule: "example-.com", wantErr: true},
		{name: "overlong label", typ: "domain", rule: strings.Repeat("a", 64) + ".com", wantErr: true},
		{name: "empty attr", typ: "domain", rule: "example.com @", wantErr: true},
		{name: "empty affiliation", typ: "domain", rule: "example.com &", wantErr: true},
		{name: "unknown field", typ: "domain", rule: "example.com ads", wantErr: true},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			entry, affs, err := parseEntry(tc.typ, tc.rule)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("parseEntry(%q, %q) = %q, want error", tc.typ, tc.rule, entry.Plain)
				}
				return
			}
			if err != nil {
				t.Fatalf("parseEntry(%q, %q) got unexpected error: %v", tc.typ, tc.rule, err)
			}
			if entry.Plain != tc.wantPlain {
				t.Errorf("parseEntry(%q, %q) = %q, want %q", tc.typ, tc.rule, entry.Plain, tc.wantPlain)
			}
			if len(affs) != len(tc.wantAffs) {
				t.Fatalf("parseEntry(%q, %q) affiliations = %v, want %v", tc.typ, tc.rule, affs, tc.wantAffs)
			}
			for i, aff := range affs {
				if aff != tc.wantAffs[i] {
					t.Errorf("parseEntry(%q, %q) affiliations = %v, want %v", tc.typ, tc.rule, affs, tc.wantAffs)
				}
			}
		})
	}
}

func TestParseInclusion(t *testing.T) {
	testCases := []struct {
		name     string
		rule     string
		wantSrc  string
		wantMust []string
		wantBan  []string
		wantErr  bool
	}{
		{name: "plain", rule: "other-list", wantSrc: "other-list"},
		{name: "filters", rule: "other @ads @-cn", wantSrc: "other", wantMust: []string{"ads"}, wantBan: []string{"cn"}},
		{name: "empty attr", rule: "other @", wantErr: true},
		{name: "empty ban attr", rule: "other @-", wantErr: true},
		{name: "empty rule", rule: " ", wantErr: true},
		{name: "invalid name", rule: "other@list", wantErr: true},
		{name: "affiliation", rule: "other &another", wantErr: true},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			inc, err := parseInclusion(tc.rule)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("parseInclusion(%q) = %+v, want error", tc.rule, inc)
				}
				return
			}
			if err != nil {
				t.Fatalf("parseInclusion(%q) got unexpected error: %v", tc.rule, err)
			}
			if inc.Source != tc.wantSrc {
				t.Errorf("parseInclusion(%q) source = %q, want %q", tc.rule, inc.Source, tc.wantSrc)
			}
			if len(inc.MustAttrs) != len(tc.wantMust) || len(inc.BanAttrs) != len(tc.wantBan) {
				t.Fatalf("parseInclusion(%q) filters = %v/%v, want %v/%v", tc.rule, inc.MustAttrs, inc.BanAttrs, tc.wantMust, tc.wantBan)
			}
			for i, attr := range inc.MustAttrs {
				if attr != tc.wantMust[i] {
					t.Errorf("parseInclusion(%q) must attrs = %v, want %v", tc.rule, inc.MustAttrs, tc.wantMust)
				}
			}
			for i, attr := range inc.BanAttrs {
				if attr != tc.wantBan[i] {
					t.Errorf("parseInclusion(%q) ban attrs = %v, want %v", tc.rule, inc.BanAttrs, tc.wantBan)
				}
			}
		})
	}
}

func TestPolishList(t *testing.T) {
	rules := []struct{ typ, rule string }{
		{"domain", "example.com @cn"},
		{"domain", "sub.example.com"},      // Redundant, no attribute
		{"full", "www.example.com @cn"},    // Redundant, same attribute
		{"full", "example.com"},            // Redundant, no attribute
		{"full", "example.org"},            // Kept, no parent domain rule
		{"domain", "ads.example.com @ads"}, // Kept, different attribute
		{"keyword", "example"},
	}
	roughMap := make(map[string]*Entry, len(rules))
	for _, r := range rules {
		entry, _, err := parseEntry(r.typ, r.rule)
		if err != nil {
			t.Fatalf("parseEntry(%q, %q) got unexpected error: %v", r.typ, r.rule, err)
		}
		roughMap[entry.Plain] = entry
	}
	want := []string{"domain:ads.example.com:@ads", "domain:example.com:@cn", "full:example.org", "keyword:example"}
	assertPlains(t, "polishList", polishList(roughMap), want)
}

// TestResolveSelectiveInclusion makes sure that selective inclusion does not
// lose rules which are redundant in the source list only.
func TestResolveSelectiveInclusion(t *testing.T) {
	dataPath := t.TempDir()
	files := map[string]string{
		"source": "domain:example.com @cn\nfull:mail.example.com\ndomain:sub.example.com\ndomain:example.org @ads\n",
		"banned": "include:source @-cn\n",
		"must":   "include:source @ads\n",
		"full":   "include:source\n",
	}
	for name, content := range files {
		if err := os.WriteFile(filepath.Join(dataPath, name), []byte(content), 0644); err != nil {
			t.Fatalf("failed to write test data %q: %v", name, err)
		}
	}

	processor := &Processor{parsedListByName: make(map[string]*ParsedList)}
	for name := range files {
		if err := processor.loadData(name, filepath.Join(dataPath, name)); err != nil {
			t.Fatalf("loadData(%q) got unexpected error: %v", name, err)
		}
	}
	for name := range files {
		if _, err := processor.resolveList(name); err != nil {
			t.Fatalf("resolveList(%q) got unexpected error: %v", name, err)
		}
	}

	assertList(t, processor, "source", []string{"domain:example.com:@cn", "domain:example.org:@ads"})
	assertList(t, processor, "full", []string{"domain:example.com:@cn", "domain:example.org:@ads"})
	assertList(t, processor, "must", []string{"domain:example.org:@ads"})
	assertList(t, processor, "banned", []string{"domain:example.org:@ads", "domain:sub.example.com", "full:mail.example.com"})
}

func TestResolveCircularInclusion(t *testing.T) {
	dataPath := t.TempDir()
	files := map[string]string{
		"first":  "domain:example.com\ninclude:second\n",
		"second": "include:first\n",
	}
	for name, content := range files {
		if err := os.WriteFile(filepath.Join(dataPath, name), []byte(content), 0644); err != nil {
			t.Fatalf("failed to write test data %q: %v", name, err)
		}
	}
	processor := &Processor{parsedListByName: make(map[string]*ParsedList)}
	for name := range files {
		if err := processor.loadData(name, filepath.Join(dataPath, name)); err != nil {
			t.Fatalf("loadData(%q) got unexpected error: %v", name, err)
		}
	}
	if _, err := processor.resolveList("first"); err == nil {
		t.Fatal("resolveList(\"first\") = nil, want circular inclusion error")
	}
}

func assertList(t *testing.T, p *Processor, name string, want []string) {
	t.Helper()
	pl, exist := p.parsedListByName[name]
	if !exist {
		t.Fatalf("list %q does not exist", name)
	}
	assertPlains(t, name, pl.FinalEntries, want)
}

func assertPlains(t *testing.T, name string, entries []*Entry, want []string) {
	t.Helper()
	got := make([]string, 0, len(entries))
	for _, entry := range entries {
		got = append(got, entry.Plain)
	}
	if !slices.Equal(got, want) {
		t.Errorf("%s = %v, want %v", name, got, want)
	}
}

func TestAssembleDat(t *testing.T) {
	outDir := t.TempDir()
	*outputDir = outDir

	siteCustom := &router.GeoSite{
		CountryCode: "CUSTOM",
		Domain: []*router.Domain{
			{Type: router.Domain_RootDomain, Value: "example.com"},
		},
	}
	siteTelegram := &router.GeoSite{
		CountryCode: "TELEGRAM",
		Domain: []*router.Domain{
			{Type: router.Domain_RootDomain, Value: "t.me"},
		},
	}

	gs := &GeoSites{
		Sites: []*router.GeoSite{siteCustom, siteTelegram},
		SiteIdx: map[string]int{
			"CUSTOM":   0,
			"TELEGRAM": 1,
		},
	}

	t.Run("ModeAll", func(t *testing.T) {
		task := DatTask{Name: "all.dat", Mode: ModeAll}
		if err := gs.assembleDat(task); err != nil {
			t.Fatalf("assembleDat failed: %v", err)
		}
		data, err := os.ReadFile(filepath.Join(outDir, "all.dat"))
		if err != nil {
			t.Fatalf("failed to read output dat: %v", err)
		}
		list := new(router.GeoSiteList)
		if err := proto.Unmarshal(data, list); err != nil {
			t.Fatalf("failed to unmarshal output: %v", err)
		}
		if len(list.Entry) != 2 {
			t.Errorf("got %d entries, want 2", len(list.Entry))
		}
	})

	t.Run("ModeAllowlist case insensitive", func(t *testing.T) {
		task := DatTask{
			Name:  "allow.dat",
			Mode:  ModeAllowlist,
			Lists: []string{"custom", "TELEGRAM"},
		}
		if err := gs.assembleDat(task); err != nil {
			t.Fatalf("assembleDat failed: %v", err)
		}
		data, err := os.ReadFile(filepath.Join(outDir, "allow.dat"))
		if err != nil {
			t.Fatalf("failed to read output dat: %v", err)
		}
		list := new(router.GeoSiteList)
		if err := proto.Unmarshal(data, list); err != nil {
			t.Fatalf("failed to unmarshal output: %v", err)
		}
		if len(list.Entry) != 2 {
			t.Fatalf("got %d entries, want 2", len(list.Entry))
		}
		if list.Entry[0].CountryCode != "CUSTOM" || list.Entry[1].CountryCode != "TELEGRAM" {
			t.Errorf("unexpected entries: %v", list.Entry)
		}
	})

	t.Run("ModeAllowlist missing list", func(t *testing.T) {
		task := DatTask{
			Name:  "missing.dat",
			Mode:  ModeAllowlist,
			Lists: []string{"nonexistent"},
		}
		if err := gs.assembleDat(task); err == nil {
			t.Fatal("assembleDat succeeded, want error for nonexistent list")
		}
	})

	t.Run("ModeDenylist case insensitive", func(t *testing.T) {
		task := DatTask{
			Name:  "deny.dat",
			Mode:  ModeDenylist,
			Lists: []string{"telegram"},
		}
		if err := gs.assembleDat(task); err != nil {
			t.Fatalf("assembleDat failed: %v", err)
		}
		data, err := os.ReadFile(filepath.Join(outDir, "deny.dat"))
		if err != nil {
			t.Fatalf("failed to read output dat: %v", err)
		}
		list := new(router.GeoSiteList)
		if err := proto.Unmarshal(data, list); err != nil {
			t.Fatalf("failed to unmarshal output: %v", err)
		}
		if len(list.Entry) != 1 {
			t.Fatalf("got %d entries, want 1", len(list.Entry))
		}
		if list.Entry[0].CountryCode != "CUSTOM" {
			t.Errorf("got entry %q, want CUSTOM", list.Entry[0].CountryCode)
		}
	})

	t.Run("ModeDenylist nothing to deny error", func(t *testing.T) {
		task := DatTask{
			Name:  "deny_none.dat",
			Mode:  ModeDenylist,
			Lists: []string{"nonexistent"},
		}
		if err := gs.assembleDat(task); err == nil {
			t.Fatal("assembleDat succeeded, want error when no lists to deny")
		}
	})
}

func TestLoadDataBOM(t *testing.T) {
	dataPath := t.TempDir()
	filePath := filepath.Join(dataPath, "bomtest")
	content := "\ufeffdomain:example.com\n"
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write test file: %v", err)
	}

	processor := &Processor{parsedListByName: make(map[string]*ParsedList)}
	if err := processor.loadData("bomtest", filePath); err != nil {
		t.Fatalf("loadData failed: %v", err)
	}
	pl, ok := processor.parsedListByName["bomtest"]
	if !ok || len(pl.Entries) != 1 {
		t.Fatalf("expected 1 entry, got %v", pl)
	}
	if pl.Entries[0].Plain != "domain:example.com" {
		t.Errorf("got %q, want domain:example.com", pl.Entries[0].Plain)
	}
}

