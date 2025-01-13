package slave

import (
	"context"
	"fmt"

	"github.com/cresplanex/bloader/internal/runner"
	"github.com/cresplanex/bloader/internal/slave/slcontainer"
)

// TmplFactor represents the slave template factor
type TmplFactor struct {
	loader                        *slcontainer.Loader
	connectionID                  string
	receiveChanelRequestContainer *slcontainer.ReceiveChanelRequestContainer
	mapper                        *slcontainer.RequestConnectionMapper
}

// TmplFactorize is a function that factorizes the template
func (s *TmplFactor) TmplFactorize(ctx context.Context, path string) (string, error) {
	l, ok := s.loader.GetLoader(path)
	if ok {
		return l, nil
	}

	term := s.receiveChanelRequestContainer.SendLoaderResourceRequests(
		ctx,
		s.connectionID,
		s.mapper,
		slcontainer.LoaderResourceRequest{
			LoaderID: path,
		},
	)
	if term == nil {
		return "", fmt.Errorf("failed to send loader resource request")
	}
	select {
	case <-ctx.Done():
		return "", nil
	case <-term:
	}

	l, ok = s.loader.GetLoader(path)
	if ok {
		return l, nil
	}

	return "", fmt.Errorf("failed to factorize the template")
}

var _ runner.TmplFactor = &TmplFactor{}
