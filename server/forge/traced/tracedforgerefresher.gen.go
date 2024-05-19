// Code generated by gowrap. DO NOT EDIT.
// template: ../../../templates/opentelemetry.go.tpl
// gowrap: http://github.com/hexdigest/gowrap

package traced

//go:generate gowrap gen -p go.woodpecker-ci.org/woodpecker/v2/server/forge/traced -i ForgeRefresher -t ../../../templates/opentelemetry.go.tpl -o tracedforgerefresher.gen.go -l ""

import (
	"context"
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"go.woodpecker-ci.org/woodpecker/v2/server/forge/types"
	"go.woodpecker-ci.org/woodpecker/v2/server/model"
)

// ForgeRefresherWithTracing implements ForgeRefresher interface instrumented with opentracing spans
type ForgeRefresherWithTracing struct {
	ForgeRefresher
	_instance      string
	_spanDecorator func(span trace.Span, params, results map[string]interface{})
}

// NewForgeRefresherWithTracing returns ForgeRefresherWithTracing
func NewForgeRefresherWithTracing(base ForgeRefresher, instance string, spanDecorator ...func(span trace.Span, params, results map[string]interface{})) ForgeRefresherWithTracing {
	d := ForgeRefresherWithTracing{
		ForgeRefresher: base,
		_instance:      instance,
	}

	if len(spanDecorator) > 0 && spanDecorator[0] != nil {
		d._spanDecorator = spanDecorator[0]
	}

	return d
}

// Activate implements ForgeRefresher
func (_d ForgeRefresherWithTracing) Activate(ctx context.Context, u *model.User, r *model.Repo, link string) (err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "ForgeRefresher.Activate")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":  ctx,
				"u":    u,
				"r":    r,
				"link": link}, map[string]interface{}{
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.ForgeRefresher.Activate(ctx, u, r, link)
}

// Auth implements ForgeRefresher
func (_d ForgeRefresherWithTracing) Auth(ctx context.Context, token string, secret string) (s1 string, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "ForgeRefresher.Auth")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":    ctx,
				"token":  token,
				"secret": secret}, map[string]interface{}{
				"s1":  s1,
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.ForgeRefresher.Auth(ctx, token, secret)
}

// BranchHead implements ForgeRefresher
func (_d ForgeRefresherWithTracing) BranchHead(ctx context.Context, u *model.User, r *model.Repo, branch string) (cp1 *model.Commit, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "ForgeRefresher.BranchHead")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":    ctx,
				"u":      u,
				"r":      r,
				"branch": branch}, map[string]interface{}{
				"cp1": cp1,
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.ForgeRefresher.BranchHead(ctx, u, r, branch)
}

// Branches implements ForgeRefresher
func (_d ForgeRefresherWithTracing) Branches(ctx context.Context, u *model.User, r *model.Repo, p *model.ListOptions) (sa1 []string, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "ForgeRefresher.Branches")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx": ctx,
				"u":   u,
				"r":   r,
				"p":   p}, map[string]interface{}{
				"sa1": sa1,
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.ForgeRefresher.Branches(ctx, u, r, p)
}

// Deactivate implements ForgeRefresher
func (_d ForgeRefresherWithTracing) Deactivate(ctx context.Context, u *model.User, r *model.Repo, link string) (err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "ForgeRefresher.Deactivate")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":  ctx,
				"u":    u,
				"r":    r,
				"link": link}, map[string]interface{}{
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.ForgeRefresher.Deactivate(ctx, u, r, link)
}

// Dir implements ForgeRefresher
func (_d ForgeRefresherWithTracing) Dir(ctx context.Context, u *model.User, r *model.Repo, b *model.Pipeline, f string) (fpa1 []*types.FileMeta, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "ForgeRefresher.Dir")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx": ctx,
				"u":   u,
				"r":   r,
				"b":   b,
				"f":   f}, map[string]interface{}{
				"fpa1": fpa1,
				"err":  err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.ForgeRefresher.Dir(ctx, u, r, b, f)
}

// File implements ForgeRefresher
func (_d ForgeRefresherWithTracing) File(ctx context.Context, u *model.User, r *model.Repo, b *model.Pipeline, f string) (ba1 []byte, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "ForgeRefresher.File")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx": ctx,
				"u":   u,
				"r":   r,
				"b":   b,
				"f":   f}, map[string]interface{}{
				"ba1": ba1,
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.ForgeRefresher.File(ctx, u, r, b, f)
}

// Hook implements ForgeRefresher
func (_d ForgeRefresherWithTracing) Hook(ctx context.Context, r *http.Request) (repo *model.Repo, pipeline *model.Pipeline, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "ForgeRefresher.Hook")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx": ctx,
				"r":   r}, map[string]interface{}{
				"repo":     repo,
				"pipeline": pipeline,
				"err":      err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.ForgeRefresher.Hook(ctx, r)
}

// Login implements ForgeRefresher
func (_d ForgeRefresherWithTracing) Login(ctx context.Context, r *types.OAuthRequest) (up1 *model.User, s1 string, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "ForgeRefresher.Login")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx": ctx,
				"r":   r}, map[string]interface{}{
				"up1": up1,
				"s1":  s1,
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.ForgeRefresher.Login(ctx, r)
}

// Org implements ForgeRefresher
func (_d ForgeRefresherWithTracing) Org(ctx context.Context, u *model.User, org string) (op1 *model.Org, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "ForgeRefresher.Org")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx": ctx,
				"u":   u,
				"org": org}, map[string]interface{}{
				"op1": op1,
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.ForgeRefresher.Org(ctx, u, org)
}

// OrgMembership implements ForgeRefresher
func (_d ForgeRefresherWithTracing) OrgMembership(ctx context.Context, u *model.User, org string) (op1 *model.OrgPerm, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "ForgeRefresher.OrgMembership")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx": ctx,
				"u":   u,
				"org": org}, map[string]interface{}{
				"op1": op1,
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.ForgeRefresher.OrgMembership(ctx, u, org)
}

// PullRequests implements ForgeRefresher
func (_d ForgeRefresherWithTracing) PullRequests(ctx context.Context, u *model.User, r *model.Repo, p *model.ListOptions) (ppa1 []*model.PullRequest, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "ForgeRefresher.PullRequests")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx": ctx,
				"u":   u,
				"r":   r,
				"p":   p}, map[string]interface{}{
				"ppa1": ppa1,
				"err":  err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.ForgeRefresher.PullRequests(ctx, u, r, p)
}

// Refresh implements ForgeRefresher
func (_d ForgeRefresherWithTracing) Refresh(ctx context.Context, up1 *model.User) (b1 bool, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "ForgeRefresher.Refresh")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx": ctx,
				"up1": up1}, map[string]interface{}{
				"b1":  b1,
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.ForgeRefresher.Refresh(ctx, up1)
}

// Repo implements ForgeRefresher
func (_d ForgeRefresherWithTracing) Repo(ctx context.Context, u *model.User, remoteID model.ForgeRemoteID, owner string, name string) (rp1 *model.Repo, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "ForgeRefresher.Repo")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":      ctx,
				"u":        u,
				"remoteID": remoteID,
				"owner":    owner,
				"name":     name}, map[string]interface{}{
				"rp1": rp1,
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.ForgeRefresher.Repo(ctx, u, remoteID, owner, name)
}

// Repos implements ForgeRefresher
func (_d ForgeRefresherWithTracing) Repos(ctx context.Context, u *model.User) (rpa1 []*model.Repo, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "ForgeRefresher.Repos")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx": ctx,
				"u":   u}, map[string]interface{}{
				"rpa1": rpa1,
				"err":  err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.ForgeRefresher.Repos(ctx, u)
}

// Status implements ForgeRefresher
func (_d ForgeRefresherWithTracing) Status(ctx context.Context, u *model.User, r *model.Repo, b *model.Pipeline, p *model.Workflow) (err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "ForgeRefresher.Status")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx": ctx,
				"u":   u,
				"r":   r,
				"b":   b,
				"p":   p}, map[string]interface{}{
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.ForgeRefresher.Status(ctx, u, r, b, p)
}

// Teams implements ForgeRefresher
func (_d ForgeRefresherWithTracing) Teams(ctx context.Context, u *model.User) (tpa1 []*model.Team, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "ForgeRefresher.Teams")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx": ctx,
				"u":   u}, map[string]interface{}{
				"tpa1": tpa1,
				"err":  err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.ForgeRefresher.Teams(ctx, u)
}
