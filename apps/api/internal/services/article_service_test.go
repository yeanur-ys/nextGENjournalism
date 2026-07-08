package services

import "testing"

func TestValidateArticleInput(t *testing.T) {
	tests := []struct {
		name    string
		input   ArticleInput
		wantErr bool
	}{
		{name: "valid draft", input: ArticleInput{Title: "Headline", Content: "Story", Status: "draft"}},
		{name: "valid published", input: ArticleInput{Title: "Headline", Content: "Story", Status: "published"}},
		{name: "missing content", input: ArticleInput{Title: "Headline", Status: "draft"}, wantErr: true},
		{name: "invalid status", input: ArticleInput{Title: "Headline", Content: "Story", Status: "bad"}, wantErr: true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validateArticleInput(tc.input)
			if tc.wantErr && err == nil {
				t.Fatal("expected error")
			}
			if !tc.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}
