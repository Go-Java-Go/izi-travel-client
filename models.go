package izi_client

import (
	"fmt"
	"strconv"
	"strings"
)

type internalRequest struct {
	endpoint string
	method   string

	withRequest     interface{}
	withResponse    interface{}
	withQueryParams map[string]string

	acceptedStatusCodes []int

	functionName string
	apiName      string
}

type ObjectType int

const (
	TypeMuseum = iota
	TypeTour
	TypeCity
	TypeCountry
	TypeCollection
	TypeExhibit
	TypeTouristAttraction
	TypeStoryNavigation
)

func (t ObjectType) String() string {
	return [...]string{"museum", "tour", "city", "country", "collection", "exhibit", "tourist_attraction", "story_navigation"}[t]
}

type formType int

const (
	compactForm = iota
	fullForm
)

func (t formType) String() string {
	return [...]string{"compact", "full"}[t]
}

type SearchBaseRequestParams struct {
	Languages []string
	Includes  []string
	Except    []string
	form      formType
}

type SearchPaginationBaseRequestParams struct {
	SearchBaseRequestParams
	Limit  int
	Offset int
}

type SearchObjectRequest struct {
	SearchPaginationBaseRequestParams
	Types  []ObjectType
	SortBy string
	Query  string
	Region string
	Lat    float64
	Lon    float64
	Radius int
}

func (s SearchBaseRequestParams) toQueryParam() map[string]string {
	q := make(map[string]string, 4)
	languages := s.Languages
	if languages != nil && len(languages) != 0 {
		q["languages"] = strings.Join(languages, ",")
	}

	includes := s.Includes
	if includes != nil && len(includes) != 0 {
		q["includes"] = strings.Join(includes, ",")
	}

	except := s.Except
	if except != nil && len(except) != 0 {
		q["except"] = strings.Join(except, ",")
	}

	if s.form != 0 {
		q["from"] = s.form.String()
	}

	return q
}

func (s SearchPaginationBaseRequestParams) toQueryParam() map[string]string {
	q := s.SearchBaseRequestParams.toQueryParam()
	if s.Limit != 0 {
		q["limit"] = strconv.Itoa(s.Limit)
	}
	if s.Offset != 0 {
		q["offset"] = strconv.Itoa(s.Offset)
	}
	return q
}

func (s SearchObjectRequest) toQueryParam() map[string]string {
	q := s.SearchPaginationBaseRequestParams.toQueryParam()

	if s.Types != nil || len(s.Types) != 0 {
		types := make([]string, len(s.Types))
		for _, el := range s.Types {
			types = append(types, el.String())
		}
		q["type"] = strings.Join(types, ",")
	}
	if s.SortBy != "" {
		q["sort_by"] = s.SortBy
	}

	if s.Lat != 0 && s.Lon != 0 && s.Radius != 0 {
		q["lat_lon"] = fmt.Sprintf("%f,%f", s.Lat, s.Lon)
		q["radius"] = strconv.Itoa(s.Radius)
	}

	return q
}

type QueryParam interface {
	toQueryParam() map[string]string
}

type CityFullForm struct {
	Uuid         string   `json:"uuid"`
	Type         string   `json:"type"`
	Languages    []string `json:"Languages"`
	Status       string   `json:"status"`
	Translations []struct {
		Name     string `json:"name"`
		Language string `json:"language"`
	} `json:"translations"`
	Map struct {
		Bounds string `json:"bounds"`
	} `json:"map"`
	Hash    string `json:"hash"`
	Visible bool   `json:"visible"`
	Content []struct {
		Title    string `json:"title"`
		Summary  string `json:"summary"`
		Desc     string `json:"desc"`
		Language string `json:"language"`
		Images   []struct {
			Uuid  string `json:"uuid"`
			Type  string `json:"type"`
			Order int    `json:"order"`
		} `json:"images"`
	} `json:"content"`
	Location Location `json:"location"`
}

type CityCompactForm struct {
	City City `json:"city"`
}

type CountryCompactForm struct {
	Country Country `json:"country"`
}

type ObjectFullForm struct {
	Uuid      string   `json:"uuid"`
	Status    string   `json:"status"`
	Type      string   `json:"type"`
	Languages []string `json:"languages"`
	Map       struct {
		Bounds string `json:"bounds"`
	} `json:"map"`
	Hash            string  `json:"hash"`
	Size            int     `json:"size"`
	City            City    `json:"city"`
	Country         Country `json:"country"`
	ContentProvider struct {
		Uuid      string `json:"uuid"`
		Name      string `json:"name"`
		Copyright string `json:"copyright"`
	} `json:"content_provider"`
	Contacts struct {
		Country string `json:"country"`
		City    string `json:"city"`
		Address string `json:"address"`
	} `json:"contacts"`
	Publisher Publisher `json:"publisher"`
	Content   []struct {
		Audio    []Audio `json:"audio"`
		Images   []Image `json:"images"`
		Download struct {
			MapMbtiles struct {
				Md5       string `json:"md5"`
				Size      int    `json:"size"`
				Url       string `json:"url"`
				UpdatedAt string `json:"updated_at"`
			} `json:"map-mbtiles"`
		} `json:"download"`
		Language string `json:"language"`
		Summary  string `json:"summary"`
		Desc     string `json:"desc"`
		Title    string `json:"title"`
	} `json:"content"`
	Location struct {
		Altitude    int     `json:"altitude"`
		CityUuid    string  `json:"city_uuid"`
		CountryCode string  `json:"country_code"`
		CountryUuid string  `json:"country_uuid"`
		Latitude    float64 `json:"latitude"`
		Longitude   float64 `json:"longitude"`
	} `json:"location"`
}

type ObjectCompactForm struct {
	Uuid      string   `json:"uuid"`
	Status    string   `json:"status"`
	Type      string   `json:"type"`
	Languages []string `json:"languages"`
	Map       struct {
		Bounds string `json:"bounds"`
	} `json:"map"`
	Hash            string  `json:"hash"`
	ChildrenCount   int     `json:"children_count"`
	City            City    `json:"city"`
	Country         Country `json:"country"`
	ContentProvider struct {
		Uuid      string `json:"uuid"`
		Name      string `json:"name"`
		Copyright string `json:"copyright"`
	} `json:"content_provider"`
	Images    []Image   `json:"images"`
	Publisher Publisher `json:"publisher"`
	Location  Location  `json:"location"`
	Language  string    `json:"language"`
	Summary   string    `json:"summary"`
	Title     string    `json:"title"`
}

type City struct {
	Uuid      string   `json:"uuid"`
	Type      string   `json:"type"`
	Languages []string `json:"languages"`
	Status    string   `json:"status"`
	Map       struct {
		Bounds string `json:"bounds"`
	} `json:"map"`
	Hash     string   `json:"hash"`
	Visible  bool     `json:"visible"`
	Title    string   `json:"title"`
	Summary  string   `json:"summary"`
	Language string   `json:"language"`
	Location Location `json:"location"`
}

type Location struct {
	Altitude    int     `json:"altitude"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	CountryCode string  `json:"country_code"`
	CountryUuid string  `json:"country_uuid"`
}

type Country struct {
	Uuid      string   `json:"uuid"`
	Type      string   `json:"type"`
	Languages []string `json:"languages"`
	Status    string   `json:"status"`
	Map       struct {
		Bounds string `json:"bounds"`
	} `json:"map"`
	Hash        string `json:"hash"`
	CountryCode string `json:"country_code"`
	Title       string `json:"title"`
	Summary     string `json:"summary"`
	Language    string `json:"language"`
	Location    struct {
		Altitude  int     `json:"altitude"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"location"`
}

type Publisher struct {
	Uuid            string   `json:"uuid"`
	Type            string   `json:"type"`
	Languages       []string `json:"languages"`
	Status          string   `json:"status"`
	Hash            string   `json:"hash"`
	Title           string   `json:"title"`
	Summary         string   `json:"summary"`
	Language        string   `json:"language"`
	Images          []Image  `json:"images"`
	ContentProvider struct {
		Uuid      string `json:"uuid"`
		Name      string `json:"name"`
		Copyright string `json:"copyright"`
	} `json:"content_provider"`
}

type Image struct {
	Uuid  string `json:"uuid"`
	Type  string `json:"type"`
	Order int    `json:"order"`
	Hash  string `json:"hash"`
	Size  int    `json:"size"`
}

type Audio struct {
	Uuid     string `json:"uuid"`
	Type     string `json:"type"`
	Duration int    `json:"duration"`
	Order    int    `json:"order"`
	Hash     string `json:"hash"`
	Size     int    `json:"size"`
}
