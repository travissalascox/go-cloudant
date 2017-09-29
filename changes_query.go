package cloudant

import (
	"net/url"
	"strconv"
)

type ChangesQueryBuilder interface {
	Conflicts()			ChangesQueryBuilder
	Descending() 		ChangesQueryBuilder
	Feed(string) 		ChangesQueryBuilder
	Filter(string) 		ChangesQueryBuilder
	Heartbeat(int)		ChangesQueryBuilder
	IncludeDocs()		ChangesQueryBuilder
	Limit(int)			ChangesQueryBuilder
	Since(string)		ChangesQueryBuilder
	Style(string)		ChangesQueryBuilder
	Timeout(int)		ChangesQueryBuilder
	Build()				QueryBuilder
}

type changesQueryBuilder struct {
	conflicts		bool
	descending		bool
	feed 			string
	filter 			string
	heartbeat		int
	include_docs	bool
	limit			int
	since			string
	style			string
	timeout			int
}

type changesQuery struct {
	conflicts		bool
	descending		bool
	feed 			string
	filter 			string
	heartbeat		int
	include_docs	bool
	limit			int
	since			string
	style			string
	timeout			int
}

func NewChangesQuery() ChangesQueryBuilder {
	return &changesQueryBuilder{}
}

func (c *changesQueryBuilder) Conflicts() ChangesQueryBuilder {
	c.conflicts = true
	return c
}

func (c *changesQueryBuilder) Descending() ChangesQueryBuilder {
	c.descending = true
	return c
}

func (c *changesQueryBuilder) Feed(feed string) ChangesQueryBuilder {
	c.feed = feed
	return c
}

func (c *changesQueryBuilder) Filter(filter string) ChangesQueryBuilder {
	c.filter = filter
	return c
}

func (c *changesQueryBuilder) Heartbeat(hb int) ChangesQueryBuilder {
	c.heartbeat = hb
	return c
}

func (c *changesQueryBuilder) IncludeDocs() ChangesQueryBuilder {
	c.include_docs = true
	return c
}

func (c *changesQueryBuilder) Limit(lim int) ChangesQueryBuilder {
	c.limit = lim
	return c
}

func (c *changesQueryBuilder) Since(seq string) ChangesQueryBuilder {
	c.since = seq
	return c
}

func (c *changesQueryBuilder) Style(style string) ChangesQueryBuilder {
	c.style = style
	return c
}

func (c *changesQueryBuilder) Timeout(secs int) ChangesQueryBuilder {
	c.timeout = secs
	return c
}

func (cq *changesQuery) QueryString() (url.Values, error) {
	vals := url.Values{}
	if cq.conflicts {
		vals.Set("conflicts", "true")
	}
	if cq.descending {
		vals.Set("descending", "true")
	}
	if cq.include_docs {
		vals.Set("include_docs", "true")
	}
	if cq.feed != "" {
		vals.Set("feed", cq.feed)
	}
	if cq.filter != "" {
		vals.Set("filter", cq.filter)
	}
	if cq.heartbeat > 0 {
		vals.Set("heartbeat", strconv.Itoa(cq.heartbeat))
	}
	if cq.style != "" {
		vals.Set("style", cq.style)
	}
	if cq.since != "" {
		vals.Set("since", cq.since)
	}
	if cq.timeout > 0 {
		vals.Set("timeout", strconv.Itoa(cq.timeout))
	}
	return vals, nil
}

func (c *changesQueryBuilder) Build() QueryBuilder {
	return &changesQuery{
		conflicts:		c.conflicts,
		descending:		c.descending,
		feed: 			c.feed,
		filter: 		c.filter,
		heartbeat:		c.heartbeat,
		include_docs:	c.include_docs,
		limit:			c.limit,
		since:			c.since,
		style:			c.style,
		timeout:		c.timeout,
	}
}
