package opt

import "github.com/algolia/algoliasearch-client-go/algolia/opt"

func ExtractAdvancedSyntaxFeatures(opts ...interface{}) *opt.AdvancedSyntaxFeaturesOption {
	for _, o := range opts {
		if v, ok := o.(opt.AdvancedSyntaxFeaturesOption); ok {
			return &v
		}
	}
	return nil
}