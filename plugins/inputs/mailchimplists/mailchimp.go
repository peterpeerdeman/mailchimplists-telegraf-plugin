package mailchimplists

import (
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
)

type MailChimp struct {
	api *ChimpAPI

	APIKey     string `toml:"api_key"`
	DaysOld    int    `toml:"days_old"`
	CampaignID string `toml:"campaign_id"`

	Log telegraf.Logger `toml:"-"`
}

func (m *MailChimp) Description() string {
	return "Gather Mailchimp Mailing List stats"
}

func (m *MailChimp) Init() error {
	m.api = NewChimpAPI(m.APIKey, m.Log)

	return nil
}

func (m *MailChimp) Gather(acc telegraf.Accumulator) error {
	lists, err := m.api.GetLists()
	if err != nil {
		return err
	}
	for _, list := range lists.Lists {
		gatherList(acc, list)
	}

	return nil
}

func gatherList(acc telegraf.Accumulator, list List) {
	tags := make(map[string]string)
	tags["id"] = list.ID
	tags["list_name"] = list.Name
	fields := map[string]interface{}{
		"member_count":                 list.Stats.MemberCount,
		"unsubscribe_count":            list.Stats.UnsubscribeCount,
		"cleaned_count":                list.Stats.CleanedCount,
		"member_count_since_send":      list.Stats.MemberCountSinceSend,
		"unsubscribe_count_since_send": list.Stats.UnsubscribeCountSinceSend,
		"cleaned_count_since_send":     list.Stats.CleanedCountSinceSend,
		"campaign_count":               list.Stats.CampaignCount,
		"campaign_last_sent":           list.Stats.CampaignLastSent,
		"merge_field_count":            list.Stats.MergeFieldCount,
		"avg_sub_rate":                 list.Stats.AvgSubRate,
		"avg_unsub_rate":               list.Stats.AvgUnsubRate,
		"target_sub_rate":              list.Stats.TargetSubRate,
		"open_rate":                    list.Stats.OpenRate,
		"click_rate":                   list.Stats.ClickRate,
		"last_sub_date":                list.Stats.LastSubDate,
		"last_unsub_date":              list.Stats.LastUnsubDate,
	}
	acc.AddFields("mailchimplists", fields, tags)
}

func init() {
	inputs.Add("mailchimplists", func() telegraf.Input {
		return &MailChimp{}
	})
}
