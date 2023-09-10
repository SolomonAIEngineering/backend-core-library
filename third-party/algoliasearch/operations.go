package algoliasearch

import "context"

type Record struct {
	ObjectID   string `json:"objectID"`
	DataObject any    `json:"data"`
}

// Send sends a profile to algolia search
func (c *Client) Send(ctx context.Context, data Record) (*string, error) {
	// instrument this operation
	txn := c.telemetrySdk.GetTraceFromContext(ctx)
	span := c.telemetrySdk.StartSegment(txn, "algolia-search-send")
	defer span.End()

	res, err := c.index.SaveObject(data)
	if err != nil {
		return nil, err
	}

	return &res.ObjectID, err
}

// Delete deletes a profile from algolia search
func (c *Client) Delete(ctx context.Context, objectId string) error {
	// instrument this operation
	txn := c.telemetrySdk.GetTraceFromContext(ctx)
	span := c.telemetrySdk.StartSegment(txn, "algolia-search-delete")
	defer span.End()

	if _, err := c.index.DeleteObject(objectId); err != nil {
		return err
	}

	return nil
}

// Update updates a profile to algolia search
func (c *Client) Update(ctx context.Context, record Record) error {
	// instrument this operation
	txn := c.telemetrySdk.GetTraceFromContext(ctx)
	span := c.telemetrySdk.StartSegment(txn, "algolia-search-update")
	defer span.End()

	// ensure record has a populated object Id
	if record.ObjectID == "" {
		return ErrMissingObjectID
	}

	if _, err := c.index.PartialUpdateObject(record); err != nil {
		return err
	}

	return nil
}
