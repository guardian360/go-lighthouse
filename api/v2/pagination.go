package v2

// Collect exhaustively fetches every page of a paginated v2 endpoint and
// returns the concatenated items. The fetch function is invoked once per page,
// in order, starting at page 1, and must return the items on that page along
// with the response metadata used to decide whether more pages remain.
//
// Callers set their own page size (and any other query parameters) inside
// fetch; Collect only drives the page number. A typical use:
//
//	objects, err := v2.Collect(func(page int) ([]v2.ScanObject, v2.APIResponseMeta, error) {
//		resp, err := api.ScannerPlatform(id).ScanObjects().PerPage(100).Page(page).Get()
//		if err != nil {
//			return nil, v2.APIResponseMeta{}, err
//		}
//		return resp.Data, resp.Meta, nil
//	})
//
// Collect is meant for callers that need the complete result set for a single
// operation: it fails atomically (an error on any page returns no items) and
// accumulates every item in memory. It is therefore unsuited to unbounded
// result sets; callers that process pages independently should page manually.
func Collect[T any](fetch func(page int) ([]T, APIResponseMeta, error)) ([]T, error) {
	var items []T
	for page := 1; ; page++ {
		pageItems, meta, err := fetch(page)
		if err != nil {
			return nil, err
		}
		items = append(items, pageItems...)

		// LastPage is zero when the endpoint returned no pagination metadata
		// (e.g. an empty or non-paginated response); treat the fetched page as
		// the only one. The empty-page check guards against a server that
		// advertises more pages than it can deliver, preventing an endless loop.
		if meta.LastPage == 0 || meta.CurrentPage >= meta.LastPage || len(pageItems) == 0 {
			return items, nil
		}
	}
}
