package azrp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type apiGetter func(url string) (ApiResponse, error)

// apiGet is used to get a single page or results from the retail price API
func apiGet(url string) (ApiResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return ApiResponse{}, nil
	}
	defer resp.Body.Close()
	ar := ApiResponse{}
	jdec := json.NewDecoder(resp.Body)
	err = jdec.Decode(&ar)
	if err != nil {
		return ApiResponse{}, err
	}
	return ar, nil
}

// apiGetAll is another version of the apiGetter function. This version checks
// the version returns all results by checking the NextResultLink field and
// keeps fetching items until the list is exhausted.
func apiGetAll(url string) (ApiResponse, error) {
	fmt.Println("Getting", url)
	r1, err := http.Get(url)
	if err != nil {
		return ApiResponse{}, err
	}
	defer r1.Body.Close()
	ar := ApiResponse{}
	jdec := json.NewDecoder(r1.Body)
	err = jdec.Decode(&ar)
	if err != nil {
		return ApiResponse{}, err
	}

	next := ar.NextPageLink

	for {
		if next == "" {
			break
		} else {
			fmt.Println("Getting", next)
			r2, err := http.Get(next)
			if err != nil {
				return ApiResponse{}, err
			}
			defer r2.Body.Close()
			arTmp := ApiResponse{}
			jdec := json.NewDecoder(r2.Body)
			err = jdec.Decode(&arTmp)
			if err != nil {
				return ApiResponse{}, nil
			}
			//Now copy the items to the original response
			ar.Items = append(ar.Items, arTmp.Items...)
			//Update next link
			next = arTmp.NextPageLink

		}
	}
	/* fixes unexpected values */
	ar.Count = uint(len(ar.Items))
	ar.NextPageLink = ""
	return ar, nil
}
