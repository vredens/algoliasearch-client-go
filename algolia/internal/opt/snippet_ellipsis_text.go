package opt

import "github.com/algolia/algoliasearch-client-go/algolia/opt"

func ExtractSnippetEllipsisText(opts ...interface{}) *opt.SnippetEllipsisTextOption {
	for _, o := range opts {
		if v, ok := o.(opt.SnippetEllipsisTextOption); ok {
			return &v
		}
	}
	return nil
}