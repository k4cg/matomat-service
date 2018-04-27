/*
General TODO:
There sees to be too much boilderplate code. Solve this more cleverly (deepen understanding of golang for better constructs!)
*/
package api

import (
	"net/http"

	"github.com/k4cg/matomat-service/maas-server/auth"
	"github.com/k4cg/matomat-service/maas-server/matomat"
)

type ItemsApiHandler struct {
	auth    auth.AuthInterface
	matomat *matomat.Matomat
}

func NewItemsApiHandler(auth auth.AuthInterface, matomat *matomat.Matomat) *ItemsApiHandler {
	return &ItemsApiHandler{auth: auth, matomat: matomat}
}

func extractItemCreateData(r *http.Request) (string, int32, error) {
	var err error

	r.ParseForm()

	itemName, err := formGet(r, FORM_KEY_NAME)
	itemCost, err := formGetInt32(r, FORM_KEY_COST)

	return itemName, itemCost, err
}

func extractItemEditData(r *http.Request) (uint32, string, int32, error) {
	var err error

	r.ParseForm()

	itemID, err := extractIDFromModelGet(r.URL.Path)
	itemName, err := formGet(r, FORM_KEY_NAME)
	itemCost, err := formGetInt32(r, FORM_KEY_COST)

	return itemID, itemName, itemCost, err
}

func (iah *ItemsApiHandler) ItemsGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(DEFAULT_HEADER_CONTENT_TYPE_KEY, DEFAULT_HEADER_CONTENT_TYPE_VALUE_JSON)
	status, response := getErrorForbiddenResponse()
	loginUserID, err := getUserIDFromContext(r)

	if err == nil {
		if err == nil && iah.matomat.IsAllowed(loginUserID, matomat.ACTION_ITEMS_LIST) {
			items, err := iah.matomat.ItemsList()
			if err == nil {
				apiItems := make([]Item, 0)
				for _, item := range items {
					apiItems = append(apiItems, newItem(item))
				}
				status, response = getResponse(http.StatusOK, apiItems)
			} else {
				status, response = getErrorResponse(http.StatusInternalServerError, err.Error())
			}
		}
	} else {
		status, response = getErrorResponse(http.StatusInternalServerError, err.Error())
	}

	w.WriteHeader(status)
	w.Write(response)
}

func (iah *ItemsApiHandler) ItemsItemidGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(DEFAULT_HEADER_CONTENT_TYPE_KEY, DEFAULT_HEADER_CONTENT_TYPE_VALUE_JSON)
	status, response := getErrorForbiddenResponse()
	loginUserID, err := getUserIDFromContext(r)

	if err == nil {
		if iah.matomat.IsAllowed(loginUserID, matomat.ACTION_ITEMS_ITEMID_GET) {
			itemID, err := extractIDFromModelGet(r.URL.Path)
			if err == nil {
				item, err := iah.matomat.ItemGet(itemID)
				if err == nil {
					status, response = getResponse(http.StatusOK, newItem(item))
				} else {
					status, response = getErrorResponse(http.StatusInternalServerError, err.Error())
				}
			} else {
				status, response = getErrorResponse(http.StatusBadRequest, err.Error())
			}
		}
	} else {
		status, response = getErrorResponse(http.StatusInternalServerError, err.Error())
	}

	w.WriteHeader(status)
	w.Write(response)
}

func (iah *ItemsApiHandler) ItemsPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(DEFAULT_HEADER_CONTENT_TYPE_KEY, DEFAULT_HEADER_CONTENT_TYPE_VALUE_JSON)
	status, response := getErrorForbiddenResponse()
	loginUserID, err := getUserIDFromContext(r)

	if err == nil {
		if iah.matomat.IsAllowed(loginUserID, matomat.ACTION_ITEMS_CREATE) {
			name, cost, err := extractItemCreateData(r)
			if err == nil {
				item, err := iah.matomat.ItemCreate(name, cost)
				if err == nil {
					status, response = getResponse(http.StatusCreated, newItem(item))
				} else {
					status, response = getErrorResponse(http.StatusInternalServerError, err.Error())
				}
			} else {
				status, response = getErrorResponse(http.StatusBadRequest, err.Error())
			}
		}
	} else {
		status, response = getErrorResponse(http.StatusInternalServerError, err.Error())
	}

	w.WriteHeader(status)
	w.Write(response)

}

func (iah *ItemsApiHandler) ItemsItemidPut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(DEFAULT_HEADER_CONTENT_TYPE_KEY, DEFAULT_HEADER_CONTENT_TYPE_VALUE_JSON)
	status, response := getErrorForbiddenResponse()
	loginUserID, err := getUserIDFromContext(r)

	if err == nil {
		if iah.matomat.IsAllowed(loginUserID, matomat.ACTION_ITEMS_ITEMID_EDIT) {
			ID, name, cost, err := extractItemEditData(r)
			if err == nil {
				item, err := iah.matomat.ItemUpdate(ID, name, cost)
				if err == nil {
					status, response = getResponse(http.StatusOK, newItem(item))
				} else {
					status, response = getErrorResponse(http.StatusInternalServerError, err.Error())
				}
			} else {
				status, response = getErrorResponse(http.StatusBadRequest, err.Error())
			}
		}
	} else {
		status, response = getErrorResponse(http.StatusInternalServerError, err.Error())
	}

	w.WriteHeader(status)
	w.Write(response)
}

func (iah *ItemsApiHandler) ItemsItemidDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(DEFAULT_HEADER_CONTENT_TYPE_KEY, DEFAULT_HEADER_CONTENT_TYPE_VALUE_JSON)
	status, response := getErrorForbiddenResponse()
	loginUserID, err := getUserIDFromContext(r)

	if err == nil {
		if iah.matomat.IsAllowed(loginUserID, matomat.ACTION_ITEMS_ITEMID_DELETE) {
			itemID, err := extractIDFromModelGet(r.URL.Path)
			if err == nil {
				item, err := iah.matomat.ItemDelete(itemID)
				if err == nil {
					status, response = getResponse(http.StatusOK, newItem(item))
				} else {
					status, response = getErrorResponse(http.StatusInternalServerError, err.Error())
				}
			} else {
				status, response = getErrorResponse(http.StatusBadRequest, err.Error())
			}
		}
	} else {
		status, response = getErrorResponse(http.StatusInternalServerError, err.Error())
	}

	w.WriteHeader(status)
	w.Write(response)
}

func (iah *ItemsApiHandler) ItemsItemidStatsGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(DEFAULT_HEADER_CONTENT_TYPE_KEY, DEFAULT_HEADER_CONTENT_TYPE_VALUE_JSON)
	w.Header().Set("Cache-Control", "max-age=30") //allow caching by the client for 30 seconds
	status, response := getErrorForbiddenResponse()
	loginUserID, err := getUserIDFromContext(r)

	if err == nil {
		if iah.matomat.IsAllowed(loginUserID, matomat.ACTION_ITEMS_ITEMID_STATS_GET) {
			itemID, err := extractIDFromModelGet(r.URL.Path)
			if err == nil {
				item, stats, err := iah.matomat.ItemGetStats(itemID)
				if err == nil {
					status, response = getResponse(http.StatusOK, newItemStats(item, stats))
				} else {
					status, response = getErrorResponse(http.StatusInternalServerError, err.Error())
				}
			} else {
				status, response = getErrorResponse(http.StatusBadRequest, err.Error())
			}
		}
	} else {
		status, response = getErrorResponse(http.StatusInternalServerError, err.Error())
	}

	w.WriteHeader(status)
	w.Write(response)
}

func (iah *ItemsApiHandler) ItemsStatsGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(DEFAULT_HEADER_CONTENT_TYPE_KEY, DEFAULT_HEADER_CONTENT_TYPE_VALUE_JSON)
	w.Header().Set("Cache-Control", "max-age=30") //allow caching by the client for 30 seconds
	status, response := getErrorForbiddenResponse()
	loginUserID, err := getUserIDFromContext(r)

	if err == nil {
		if iah.matomat.IsAllowed(loginUserID, matomat.ACTION_ITEMS_STATS_GET) {
			items, err := iah.matomat.ItemsList()
			if err == nil {
				itemStatsList, err := iah.matomat.ItemStatsList()
				if err == nil {
					apiItemStats := make([]ItemStats, 0)
					for _, item := range items {
						itemStats, found := getStatsForItem(itemStatsList, item.ID)
						if found {
							apiItemStats = append(apiItemStats, newItemStats(item, itemStats))
						}
					}
					status, response = getResponse(http.StatusOK, apiItemStats)
				} else {
					status, response = getErrorResponse(http.StatusInternalServerError, err.Error())
				}
			} else {
				status, response = getErrorResponse(http.StatusInternalServerError, err.Error())
			}
		}
	} else {
		status, response = getErrorResponse(http.StatusInternalServerError, err.Error())
	}

	w.WriteHeader(status)
	w.Write(response)
}

func (iah *ItemsApiHandler) ItemsItemidConsumePut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(DEFAULT_HEADER_CONTENT_TYPE_KEY, DEFAULT_HEADER_CONTENT_TYPE_VALUE_JSON)
	status, response := getErrorForbiddenResponse()
	loginUserID, err := getUserIDFromContext(r)

	if err == nil {
		if iah.matomat.IsAllowed(loginUserID, matomat.ACTION_ITEMS_ITEMID_CONSUME) {
			itemID, err := extractIDFromModelGet(r.URL.Path)
			if err == nil {
				item, itemStats, err := iah.matomat.ItemConsume(loginUserID, itemID)
				if err == nil {
					status, response = getResponse(http.StatusOK, newItemStats(item, itemStats))
				} else {
					status, response = getErrorResponse(http.StatusInternalServerError, err.Error())
				}
			} else {
				status, response = getErrorResponse(http.StatusBadRequest, err.Error())
			}
		}
	} else {
		status, response = getErrorResponse(http.StatusInternalServerError, err.Error())
	}

	w.WriteHeader(status)
	w.Write(response)
}
