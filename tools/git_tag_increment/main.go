package main

func main() {
	//version, err := semver.Canonical("h1.1.1")
	//fmt.Println(version, err)
}

type T struct {
	Homebrew struct {
		Properties struct {
			Name struct {
				Type string `json:"type"`
			} `json:"name"`
			Repository struct {
				Ref string `json:"$ref"`
			} `json:"repository"`
			CommitAuthor struct {
				Ref string `json:"$ref"`
			} `json:"commit_author"`
			CommitMsgTemplate struct {
				Type string `json:"type"`
			} `json:"commit_msg_template"`
			Directory struct {
				Type string `json:"type"`
			} `json:"directory"`
			Caveats struct {
				Type string `json:"type"`
			} `json:"caveats"`
			Install struct {
				Type string `json:"type"`
			} `json:"install"`
			ExtraInstall struct {
				Type string `json:"type"`
			} `json:"extra_install"`
			PostInstall struct {
				Type string `json:"type"`
			} `json:"post_install"`
			Dependencies struct {
				Items struct {
					Ref string `json:"$ref"`
				} `json:"items"`
				Type string `json:"type"`
			} `json:"dependencies"`
			Test struct {
				Type string `json:"type"`
			} `json:"test"`
			Conflicts struct {
				Items struct {
					Type string `json:"type"`
				} `json:"items"`
				Type string `json:"type"`
			} `json:"conflicts"`
			Description struct {
				Type string `json:"type"`
			} `json:"description"`
			Homepage struct {
				Type string `json:"type"`
			} `json:"homepage"`
			License struct {
				Type string `json:"type"`
			} `json:"license"`
			SkipUpload struct {
				OneOf []struct {
					Type string `json:"type"`
				} `json:"oneOf"`
			} `json:"skip_upload"`
			DownloadStrategy struct {
				Type string `json:"type"`
			} `json:"download_strategy"`
			UrlTemplate struct {
				Type string `json:"type"`
			} `json:"url_template"`
			UrlHeaders struct {
				Items struct {
					Type string `json:"type"`
				} `json:"items"`
				Type string `json:"type"`
			} `json:"url_headers"`
			CustomRequire struct {
				Type string `json:"type"`
			} `json:"custom_require"`
			CustomBlock struct {
				Type string `json:"type"`
			} `json:"custom_block"`
			Ids struct {
				Items struct {
					Type string `json:"type"`
				} `json:"items"`
				Type string `json:"type"`
			} `json:"ids"`
			Goarm struct {
				OneOf []struct {
					Type string `json:"type"`
				} `json:"oneOf"`
			} `json:"goarm"`
			Goamd64 struct {
				Type string `json:"type"`
			} `json:"goamd64"`
			Service struct {
				Type string `json:"type"`
			} `json:"service"`
		} `json:"properties"`
		AdditionalProperties bool   `json:"additionalProperties"`
		Type                 string `json:"type"`
	} `json:"Homebrew"`
	HomebrewDependency struct {
		OneOf []struct {
			Type       string `json:"type"`
			Schema     string `json:"$schema,omitempty"`
			Id         string `json:"$id,omitempty"`
			Properties struct {
				Name struct {
					Type string `json:"type"`
				} `json:"name"`
				Type struct {
					Type string `json:"type"`
				} `json:"type"`
				Version struct {
					Type string `json:"type"`
				} `json:"version"`
				Os struct {
					Type string   `json:"type"`
					Enum []string `json:"enum"`
				} `json:"os"`
			} `json:"properties,omitempty"`
			AdditionalProperties bool `json:"additionalProperties,omitempty"`
		} `json:"oneOf"`
	} `json:"HomebrewDependency"`
}
