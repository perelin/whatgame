package microsoftgp

import "time"

type GamepassResponseAllGames []struct {
	SiglID            string `json:"siglId,omitempty"`
	Title             string `json:"title,omitempty"`
	Description       string `json:"description,omitempty"`
	RequiresShuffling string `json:"requiresShuffling,omitempty"`
	ImageURL          string `json:"imageUrl,omitempty"`
	ID                string `json:"id,omitempty"`
}

type GamepassResponseGamesDetails struct {
	Products []GamepassGameDetails `json:"Products"`
}
type GamepassGameDetails struct {
	LastModifiedDate    time.Time `json:"LastModifiedDate"`
	LocalizedProperties []struct {
		DeveloperName         string `json:"DeveloperName"`
		PublisherName         string `json:"PublisherName"`
		PublisherWebsiteURI   string `json:"PublisherWebsiteUri"`
		SupportURI            string `json:"SupportUri"`
		EligibilityProperties struct {
			Remediations []struct {
				RemediationID string `json:"RemediationId"`
				Description   string `json:"Description"`
			} `json:"Remediations"`
			Affirmations []struct {
				AffirmationID        string `json:"AffirmationId"`
				AffirmationProductID string `json:"AffirmationProductId"`
				Description          string `json:"Description"`
			} `json:"Affirmations"`
		} `json:"EligibilityProperties"`
		Franchises []interface{} `json:"Franchises"`
		Images     []struct {
			FileID                  string      `json:"FileId"`
			EISListingIdentifier    interface{} `json:"EISListingIdentifier"`
			BackgroundColor         string      `json:"BackgroundColor"`
			Caption                 string      `json:"Caption"`
			FileSizeInBytes         int         `json:"FileSizeInBytes"`
			ForegroundColor         string      `json:"ForegroundColor"`
			Height                  int         `json:"Height"`
			ImagePositionInfo       string      `json:"ImagePositionInfo"`
			ImagePurpose            string      `json:"ImagePurpose"`
			UnscaledImageSHA256Hash string      `json:"UnscaledImageSHA256Hash"`
			URI                     string      `json:"Uri"`
			Width                   int         `json:"Width"`
		} `json:"Images"`
		Videos []struct {
			URI               string `json:"Uri"`
			VideoPurpose      string `json:"VideoPurpose"`
			Height            int    `json:"Height"`
			Width             int    `json:"Width"`
			AudioEncoding     string `json:"AudioEncoding"`
			VideoEncoding     string `json:"VideoEncoding"`
			VideoPositionInfo string `json:"VideoPositionInfo"`
			Caption           string `json:"Caption"`
			FileSizeInBytes   int    `json:"FileSizeInBytes"`
			PreviewImage      struct {
				FileID                  string      `json:"FileId"`
				EISListingIdentifier    interface{} `json:"EISListingIdentifier"`
				BackgroundColor         interface{} `json:"BackgroundColor"`
				Caption                 string      `json:"Caption"`
				FileSizeInBytes         int         `json:"FileSizeInBytes"`
				ForegroundColor         interface{} `json:"ForegroundColor"`
				Height                  int         `json:"Height"`
				ImagePositionInfo       interface{} `json:"ImagePositionInfo"`
				ImagePurpose            string      `json:"ImagePurpose"`
				UnscaledImageSHA256Hash string      `json:"UnscaledImageSHA256Hash"`
				URI                     string      `json:"Uri"`
				Width                   int         `json:"Width"`
			} `json:"PreviewImage"`
			SortOrder int `json:"SortOrder"`
		} `json:"Videos"`
		ProductDescription string      `json:"ProductDescription"`
		ProductTitle       string      `json:"ProductTitle"`
		ShortTitle         string      `json:"ShortTitle"`
		SortTitle          string      `json:"SortTitle"`
		FriendlyTitle      interface{} `json:"FriendlyTitle"`
		ShortDescription   string      `json:"ShortDescription"`
		SearchTitles       []struct {
			SearchTitleString string `json:"SearchTitleString"`
			SearchTitleType   string `json:"SearchTitleType"`
		} `json:"SearchTitles"`
		VoiceTitle             string        `json:"VoiceTitle"`
		RenderGroupDetails     interface{}   `json:"RenderGroupDetails"`
		ProductDisplayRanks    []interface{} `json:"ProductDisplayRanks"`
		InteractiveModelConfig interface{}   `json:"InteractiveModelConfig"`
		Interactive3DEnabled   bool          `json:"Interactive3DEnabled"`
		Language               string        `json:"Language"`
		Markets                []string      `json:"Markets"`
	} `json:"LocalizedProperties"`
	MarketProperties []struct {
		OriginalReleaseDate time.Time `json:"OriginalReleaseDate"`
		MinimumUserAge      int       `json:"MinimumUserAge"`
		ContentRatings      []struct {
			RatingSystem        string        `json:"RatingSystem"`
			RatingID            string        `json:"RatingId"`
			RatingDescriptors   []string      `json:"RatingDescriptors"`
			RatingDisclaimers   []interface{} `json:"RatingDisclaimers"`
			InteractiveElements []interface{} `json:"InteractiveElements"`
		} `json:"ContentRatings"`
		RelatedProducts []struct {
			RelatedProductID string `json:"RelatedProductId"`
			RelationshipType string `json:"RelationshipType"`
		} `json:"RelatedProducts"`
		UsageData []struct {
			AggregateTimeSpan string      `json:"AggregateTimeSpan"`
			AverageRating     float64     `json:"AverageRating"`
			PlayCount         interface{} `json:"PlayCount"`
			RatingCount       int         `json:"RatingCount"`
			RentalCount       string      `json:"RentalCount"`
			TrialCount        string      `json:"TrialCount"`
			PurchaseCount     string      `json:"PurchaseCount"`
		} `json:"UsageData"`
		BundleConfig interface{} `json:"BundleConfig"`
		Markets      []string    `json:"Markets"`
	} `json:"MarketProperties"`
	ProductASchema string `json:"ProductASchema"`
	ProductBSchema string `json:"ProductBSchema"`
	ProductID      string `json:"ProductId"`
	Properties     struct {
		Attributes []struct {
			Name                string      `json:"Name"`
			Minimum             int         `json:"Minimum"`
			Maximum             int         `json:"Maximum"`
			ApplicablePlatforms interface{} `json:"ApplicablePlatforms"`
			Group               interface{} `json:"Group"`
		} `json:"Attributes"`
		CanInstallToSDCard                   bool        `json:"CanInstallToSDCard"`
		Category                             string      `json:"Category"`
		Categories                           []string    `json:"Categories"`
		Subcategory                          interface{} `json:"Subcategory"`
		IsAccessible                         bool        `json:"IsAccessible"`
		IsDemo                               bool        `json:"IsDemo"`
		IsLineOfBusinessApp                  bool        `json:"IsLineOfBusinessApp"`
		IsPublishedToLegacyWindowsPhoneStore bool        `json:"IsPublishedToLegacyWindowsPhoneStore"`
		IsPublishedToLegacyWindowsStore      bool        `json:"IsPublishedToLegacyWindowsStore"`
		PackageFamilyName                    string      `json:"PackageFamilyName"`
		PackageIdentityName                  string      `json:"PackageIdentityName"`
		PublisherCertificateName             string      `json:"PublisherCertificateName"`
		PublisherID                          string      `json:"PublisherId"`
		SkuDisplayGroups                     []struct {
			ID        string `json:"Id"`
			Treatment string `json:"Treatment"`
		} `json:"SkuDisplayGroups"`
		XboxLiveTier             string      `json:"XboxLiveTier"`
		XboxXPA                  interface{} `json:"XboxXPA"`
		XboxCrossGenSetID        interface{} `json:"XboxCrossGenSetId"`
		XboxConsoleGenOptimized  []string    `json:"XboxConsoleGenOptimized"`
		XboxConsoleGenCompatible []string    `json:"XboxConsoleGenCompatible"`
		XboxLiveGoldRequired     bool        `json:"XboxLiveGoldRequired"`
		ExtendedMetadata         string      `json:"ExtendedMetadata"`
		Xbox                     interface{} `json:"XBOX"`
		OwnershipType            interface{} `json:"OwnershipType"`
		PdpBackgroundColor       string      `json:"PdpBackgroundColor"`
		HasAddOns                bool        `json:"HasAddOns"`
		RevisionID               time.Time   `json:"RevisionId"`
		ProductGroupID           string      `json:"ProductGroupId"`
		ProductGroupName         string      `json:"ProductGroupName"`
	} `json:"Properties"`
	AlternateIds []struct {
		IDType string `json:"IdType"`
		Value  string `json:"Value"`
	} `json:"AlternateIds"`
	DomainDataVersion  interface{} `json:"DomainDataVersion"`
	IngestionSource    string      `json:"IngestionSource"`
	IsMicrosoftProduct bool        `json:"IsMicrosoftProduct"`
	PreferredSkuID     string      `json:"PreferredSkuId"`
	ProductType        string      `json:"ProductType"`
	// ValidationData     struct {
	// 	PassedValidation    bool      `json:"PassedValidation"`
	// 	RevisionID          time.Time `json:"RevisionId"`
	// 	ValidationResultURI string    `json:"ValidationResultUri"`
	// } `json:"ValidationData"`
	MerchandizingTags  []interface{} `json:"MerchandizingTags"`
	PartD              string        `json:"PartD"`
	SandboxID          string        `json:"SandboxId"`
	ProductFamily      string        `json:"ProductFamily"`
	SchemaVersion      string        `json:"SchemaVersion"`
	IsSandboxedProduct bool          `json:"IsSandboxedProduct"`
	ProductKind        string        `json:"ProductKind"`
	ProductPolicies    struct {
	} `json:"ProductPolicies"`
	DisplaySkuAvailabilities []struct {
		Sku struct {
			LastModifiedDate    time.Time `json:"LastModifiedDate"`
			LocalizedProperties []struct {
				Contributors              []interface{} `json:"Contributors"`
				Features                  []interface{} `json:"Features"`
				MinimumNotes              string        `json:"MinimumNotes"`
				RecommendedNotes          string        `json:"RecommendedNotes"`
				ReleaseNotes              string        `json:"ReleaseNotes"`
				DisplayPlatformProperties interface{}   `json:"DisplayPlatformProperties"`
				SkuDescription            string        `json:"SkuDescription"`
				SkuTitle                  string        `json:"SkuTitle"`
				SkuButtonTitle            string        `json:"SkuButtonTitle"`
				DeliveryDateOverlay       interface{}   `json:"DeliveryDateOverlay"`
				SkuDisplayRank            []interface{} `json:"SkuDisplayRank"`
				TextResources             interface{}   `json:"TextResources"`
				Images                    []interface{} `json:"Images"`
				LegalText                 struct {
					AdditionalLicenseTerms string `json:"AdditionalLicenseTerms"`
					Copyright              string `json:"Copyright"`
					CopyrightURI           string `json:"CopyrightUri"`
					PrivacyPolicy          string `json:"PrivacyPolicy"`
					PrivacyPolicyURI       string `json:"PrivacyPolicyUri"`
					Tou                    string `json:"Tou"`
					TouURI                 string `json:"TouUri"`
				} `json:"LegalText"`
				Language string   `json:"Language"`
				Markets  []string `json:"Markets"`
			} `json:"LocalizedProperties"`
			MarketProperties []struct {
				FirstAvailableDate time.Time   `json:"FirstAvailableDate"`
				SupportedLanguages []string    `json:"SupportedLanguages"`
				PackageIds         interface{} `json:"PackageIds"`
				PIFilter           interface{} `json:"PIFilter"`
				Markets            []string    `json:"Markets"`
			} `json:"MarketProperties"`
			ProductID  string `json:"ProductId"`
			Properties struct {
				EarlyAdopterEnrollmentURL interface{} `json:"EarlyAdopterEnrollmentUrl"`
				FulfillmentData           struct {
					ProductID         string      `json:"ProductId"`
					WuBundleID        string      `json:"WuBundleId"`
					WuCategoryID      string      `json:"WuCategoryId"`
					PackageFamilyName string      `json:"PackageFamilyName"`
					SkuID             string      `json:"SkuId"`
					Content           interface{} `json:"Content"`
					PackageFeatures   interface{} `json:"PackageFeatures"`
				} `json:"FulfillmentData"`
				FulfillmentType     string      `json:"FulfillmentType"`
				FulfillmentPluginID interface{} `json:"FulfillmentPluginId"`
				HasThirdPartyIAPs   bool        `json:"HasThirdPartyIAPs"`
				LastUpdateDate      time.Time   `json:"LastUpdateDate"`
				HardwareProperties  struct {
					MinimumHardware      []interface{} `json:"MinimumHardware"`
					RecommendedHardware  []interface{} `json:"RecommendedHardware"`
					MinimumProcessor     interface{}   `json:"MinimumProcessor"`
					RecommendedProcessor interface{}   `json:"RecommendedProcessor"`
					MinimumGraphics      interface{}   `json:"MinimumGraphics"`
					RecommendedGraphics  interface{}   `json:"RecommendedGraphics"`
				} `json:"HardwareProperties"`
				HardwareRequirements []interface{} `json:"HardwareRequirements"`
				HardwareWarningList  []interface{} `json:"HardwareWarningList"`
				InstallationTerms    string        `json:"InstallationTerms"`
				Packages             []struct {
					Applications                []interface{} `json:"Applications"`
					Architectures               []string      `json:"Architectures"`
					Capabilities                []interface{} `json:"Capabilities"`
					DeviceCapabilities          []interface{} `json:"DeviceCapabilities"`
					ExperienceIds               []interface{} `json:"ExperienceIds"`
					FrameworkDependencies       []interface{} `json:"FrameworkDependencies"`
					HardwareDependencies        []interface{} `json:"HardwareDependencies"`
					HardwareRequirements        []interface{} `json:"HardwareRequirements"`
					Hash                        string        `json:"Hash"`
					HashAlgorithm               string        `json:"HashAlgorithm"`
					IsStreamingApp              bool          `json:"IsStreamingApp"`
					Languages                   []interface{} `json:"Languages"`
					MaxDownloadSizeInBytes      int64         `json:"MaxDownloadSizeInBytes"`
					MaxInstallSizeInBytes       int           `json:"MaxInstallSizeInBytes"`
					PackageFormat               string        `json:"PackageFormat"`
					PackageFamilyName           interface{}   `json:"PackageFamilyName"`
					MainPackageFamilyNameForDlc interface{}   `json:"MainPackageFamilyNameForDlc"`
					PackageFullName             string        `json:"PackageFullName"`
					PackageID                   string        `json:"PackageId"`
					ContentID                   string        `json:"ContentId"`
					KeyID                       string        `json:"KeyId"`
					PackageRank                 int           `json:"PackageRank"`
					PackageURI                  string        `json:"PackageUri"`
					PlatformDependencies        []struct {
						MaxTested    int    `json:"MaxTested"`
						MinVersion   int    `json:"MinVersion"`
						PlatformName string `json:"PlatformName"`
					} `json:"PlatformDependencies"`
					PlatformDependencyXMLBlob string `json:"PlatformDependencyXmlBlob"`
					ResourceID                string `json:"ResourceId"`
					Version                   string `json:"Version"`
					PackageDownloadUris       []struct {
						Rank int    `json:"Rank"`
						URI  string `json:"Uri"`
					} `json:"PackageDownloadUris"`
					DriverDependencies []interface{} `json:"DriverDependencies"`
					FulfillmentData    struct {
						ProductID         string      `json:"ProductId"`
						WuBundleID        string      `json:"WuBundleId"`
						WuCategoryID      string      `json:"WuCategoryId"`
						PackageFamilyName string      `json:"PackageFamilyName"`
						SkuID             string      `json:"SkuId"`
						Content           interface{} `json:"Content"`
						PackageFeatures   interface{} `json:"PackageFeatures"`
					} `json:"FulfillmentData"`
				} `json:"Packages"`
				VersionString                 string        `json:"VersionString"`
				SkuDisplayGroupIds            []string      `json:"SkuDisplayGroupIds"`
				XboxXPA                       bool          `json:"XboxXPA"`
				BundledSkus                   []interface{} `json:"BundledSkus"`
				IsRepurchasable               bool          `json:"IsRepurchasable"`
				SkuDisplayRank                int           `json:"SkuDisplayRank"`
				DisplayPhysicalStoreInventory interface{}   `json:"DisplayPhysicalStoreInventory"`
				VisibleToB2BServiceIds        []interface{} `json:"VisibleToB2BServiceIds"`
				AdditionalIdentifiers         []interface{} `json:"AdditionalIdentifiers"`
				IsTrial                       bool          `json:"IsTrial"`
				IsPreOrder                    bool          `json:"IsPreOrder"`
				IsBundle                      bool          `json:"IsBundle"`
			} `json:"Properties"`
			SkuASchema           string      `json:"SkuASchema"`
			SkuBSchema           string      `json:"SkuBSchema"`
			SkuID                string      `json:"SkuId"`
			SkuType              string      `json:"SkuType"`
			RecurrencePolicy     interface{} `json:"RecurrencePolicy"`
			SubscriptionPolicyID interface{} `json:"SubscriptionPolicyId"`
		} `json:"Sku"`
	} `json:"DisplaySkuAvailabilities"`
}
