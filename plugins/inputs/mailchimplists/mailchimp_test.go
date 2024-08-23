package mailchimplists

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/testutil"
	"github.com/stretchr/testify/require"
)

func TestMailChimpGatherLists(t *testing.T) {
	ts := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				_, err := fmt.Fprintln(w, sampleLists)
				require.NoError(t, err)
			},
		))
	defer ts.Close()

	u, err := url.ParseRequestURI(ts.URL)
	require.NoError(t, err)

	api := &ChimpAPI{
		url:   u,
		debug: true,
		log:   testutil.Logger{},
	}
	m := MailChimp{
		api: api,
	}

	var acc testutil.Accumulator

	err = m.Gather(&acc)
	require.NoError(t, err)

	expected := []telegraf.Metric{
		testutil.MustMetric(
			"mailchimplists",
			map[string]string{
				"id":        "xxxxxdc419",
				"list_name": "Chimplist",
			},
			map[string]interface{}{
				"member_count":                 int32(23),
				"unsubscribe_count":            0,
				"cleaned_count":                1,
				"member_count_since_send":      4,
				"unsubscribe_count_since_send": 0,
				"cleaned_count_since_send":     1,
				"campaign_count":               6,
				"campaign_last_sent":           "2022-03-01T16:04:42+00:00",
				"merge_field_count":            5,
				"avg_sub_rate":                 2,
				"avg_unsub_rate":               0,
				"target_sub_rate":              0,
				"open_rate":                    float64(35.483870967741936),
				"click_rate":                   float64(13.333333333333334),
				"last_sub_date":                "2022-05-13T08:17:27+00:00",
				"last_unsub_date":              "",
			},
			time.Unix(0, 0),
		),
	}

	testutil.RequireMetricsEqual(t, expected, acc.GetTelegrafMetrics(), testutil.IgnoreTime())

	// acc.AssertContainsTaggedFields(t, "mailchimplists", fields, tags)
}

var sampleLists = `
{
  "lists": [
    {
      "id": "e2b0bdc419",
      "web_id": 1636406,
      "name": "Snoffeecob",
      "contact": {
        "company": "Acme",
        "address1": "Teststreet 1",
        "address2": "",
        "city": "Amsterdam",
        "state": "Noord-Holland",
        "zip": "1012 AB",
        "country": "NL",
        "phone": ""
      },
      "permission_reminder": "You are receiving this email because you opted in via our website.",
      "use_archive_bar": true,
      "campaign_defaults": {
        "from_name": "Test",
        "from_email": "test@acme.com",
        "subject": "",
        "language": "en"
      },
      "notify_on_subscribe": "",
      "notify_on_unsubscribe": "",
      "date_created": "2021-02-20T14:40:53+00:00",
      "list_rating": 0,
      "email_type_option": false,
      "subscribe_url_short": "http://eepurl.com/hq83KX",
      "subscribe_url_long": "https://xxx.us1.list-manage.com/subscribe?u=dd92fad27862c2ea7730a249c&id=e2b0bdc419",
      "beamer_address": "us1-6f270cd2d5-4cfab5635a@inbound.mailchimp.com",
      "visibility": "prv",
      "double_optin": false,
      "has_welcome": false,
      "marketing_permissions": false,
      "modules": [],
      "stats": {
        "member_count": 23,
        "unsubscribe_count": 0,
        "cleaned_count": 1,
        "member_count_since_send": 4,
        "unsubscribe_count_since_send": 0,
        "cleaned_count_since_send": 1,
        "campaign_count": 6,
        "campaign_last_sent": "2022-03-01T16:04:42+00:00",
        "merge_field_count": 5,
        "avg_sub_rate": 2,
        "avg_unsub_rate": 0,
        "target_sub_rate": 0,
        "open_rate": 35.483870967741936,
        "click_rate": 13.333333333333334,
        "last_sub_date": "2022-05-13T08:17:27+00:00",
        "last_unsub_date": ""
      },
      "_links": [
        {
          "rel": "self",
          "href": "https://us1.api.mailchimp.com/3.0/lists/e2b0bdc419",
          "method": "GET",
          "targetSchema": "https://us1.api.mailchimp.com/schema/3.0/Definitions/Lists/Response.json"
        },
        {
          "rel": "parent",
          "href": "https://us1.api.mailchimp.com/3.0/lists",
          "method": "GET",
          "targetSchema": "https://us1.api.mailchimp.com/schema/3.0/Definitions/Lists/CollectionResponse.json",
          "schema": "https://us1.api.mailchimp.com/schema/3.0/Paths/Lists/Collection.json"
        },
        {
          "rel": "update",
          "href": "https://us1.api.mailchimp.com/3.0/lists/e2b0bdc419",
          "method": "PATCH",
          "targetSchema": "https://us1.api.mailchimp.com/schema/3.0/Definitions/Lists/Response.json",
          "schema": "https://us1.api.mailchimp.com/schema/3.0/Definitions/Lists/PATCH.json"
        },
        {
          "rel": "batch-sub-unsub-members",
          "href": "https://us1.api.mailchimp.com/3.0/lists/e2b0bdc419",
          "method": "POST",
          "targetSchema": "https://us1.api.mailchimp.com/schema/3.0/Definitions/Lists/BatchPOST-Response.json",
          "schema": "https://us1.api.mailchimp.com/schema/3.0/Definitions/Lists/BatchPOST.json"
        },
        {
          "rel": "delete",
          "href": "https://us1.api.mailchimp.com/3.0/lists/e2b0bdc419",
          "method": "DELETE"
        },
        {
          "rel": "abuse-reports",
          "href": "https://us1.api.mailchimp.com/3.0/lists/e2b0bdc419/abuse-reports",
          "method": "GET",
          "targetSchema": "https://us1.api.mailchimp.com/schema/3.0/Definitions/Lists/Abuse/CollectionResponse.json",
          "schema": "https://us1.api.mailchimp.com/schema/3.0/Paths/Lists/Abuse/Collection.json"
        },
        {
          "rel": "activity",
          "href": "https://us1.api.mailchimp.com/3.0/lists/e2b0bdc419/activity",
          "method": "GET",
          "targetSchema": "https://us1.api.mailchimp.com/schema/3.0/Definitions/Lists/Activity/Response.json"
        },
        {
          "rel": "clients",
          "href": "https://us1.api.mailchimp.com/3.0/lists/e2b0bdc419/clients",
          "method": "GET",
          "targetSchema": "https://us1.api.mailchimp.com/schema/3.0/Definitions/Lists/Clients/Response.json"
        },
        {
          "rel": "growth-history",
          "href": "https://us1.api.mailchimp.com/3.0/lists/e2b0bdc419/growth-history",
          "method": "GET",
          "targetSchema": "https://us1.api.mailchimp.com/schema/3.0/Definitions/Lists/Growth/CollectionResponse.json",
          "schema": "https://us1.api.mailchimp.com/schema/3.0/Paths/Lists/Growth/Collection.json"
        },
        {
          "rel": "interest-categories",
          "href": "https://us1.api.mailchimp.com/3.0/lists/e2b0bdc419/interest-categories",
          "method": "GET",
          "targetSchema": "https://us1.api.mailchimp.com/schema/3.0/Definitions/Lists/InterestCategories/CollectionResponse.json",
          "schema": "https://us1.api.mailchimp.com/schema/3.0/Paths/Lists/InterestCategories/Collection.json"
        },
        {
          "rel": "members",
          "href": "https://us1.api.mailchimp.com/3.0/lists/e2b0bdc419/members",
          "method": "GET",
          "targetSchema": "https://us1.api.mailchimp.com/schema/3.0/Definitions/Lists/Members/CollectionResponse.json",
          "schema": "https://us1.api.mailchimp.com/schema/3.0/Paths/Lists/Members/Collection.json"
        },
        {
          "rel": "merge-fields",
          "href": "https://us1.api.mailchimp.com/3.0/lists/e2b0bdc419/merge-fields",
          "method": "GET",
          "targetSchema": "https://us1.api.mailchimp.com/schema/3.0/Definitions/Lists/MergeFields/CollectionResponse.json",
          "schema": "https://us1.api.mailchimp.com/schema/3.0/Paths/Lists/MergeFields/Collection.json"
        },
        {
          "rel": "segments",
          "href": "https://us1.api.mailchimp.com/3.0/lists/e2b0bdc419/segments",
          "method": "GET",
          "targetSchema": "https://us1.api.mailchimp.com/schema/3.0/Definitions/Lists/Segments/CollectionResponse.json",
          "schema": "https://us1.api.mailchimp.com/schema/3.0/Paths/Lists/Segments/Collection.json"
        },
        {
          "rel": "webhooks",
          "href": "https://us1.api.mailchimp.com/3.0/lists/e2b0bdc419/webhooks",
          "method": "GET",
          "targetSchema": "https://us1.api.mailchimp.com/schema/3.0/Definitions/Lists/Webhooks/CollectionResponse.json",
          "schema": "https://us1.api.mailchimp.com/schema/3.0/Paths/Lists/Webhooks/Collection.json"
        },
        {
          "rel": "signup-forms",
          "href": "https://us1.api.mailchimp.com/3.0/lists/e2b0bdc419/signup-forms",
          "method": "GET",
          "targetSchema": "https://us1.api.mailchimp.com/schema/3.0/Definitions/Lists/SignupForms/CollectionResponse.json",
          "schema": "https://us1.api.mailchimp.com/schema/3.0/Paths/Lists/SignupForms/Collection.json"
        },
        {
          "rel": "locations",
          "href": "https://us1.api.mailchimp.com/3.0/lists/e2b0bdc419/locations",
          "method": "GET",
          "targetSchema": "https://us1.api.mailchimp.com/schema/3.0/Definitions/Lists/Locations/CollectionResponse.json",
          "schema": "https://us1.api.mailchimp.com/schema/3.0/Paths/Lists/Locations/Collection.json"
        }
      ]
    }
  ],
  "total_items": 1,
  "constraints": {
    "may_create": false,
    "max_instances": 1,
    "current_total_instances": 1
  },
  "_links": [
    {
      "rel": "self",
      "href": "https://us1.api.mailchimp.com/3.0/lists",
      "method": "GET",
      "targetSchema": "https://us1.api.mailchimp.com/schema/3.0/Definitions/Lists/CollectionResponse.json",
      "schema": "https://us1.api.mailchimp.com/schema/3.0/Paths/Lists/Collection.json"
    },
    {
      "rel": "create",
      "href": "https://us1.api.mailchimp.com/3.0/lists",
      "method": "POST",
      "targetSchema": "https://us1.api.mailchimp.com/schema/3.0/Definitions/Lists/Response.json",
      "schema": "https://us1.api.mailchimp.com/schema/3.0/Definitions/Lists/POST.json"
    }
  ]
}
`
