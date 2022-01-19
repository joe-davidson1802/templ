// Code generated by templ@(devel) DO NOT EDIT.

package testif

import "github.com/a-h/templ"
import "context"
import "io"

func render(d data) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		ctx, _ = templ.RenderedCSSClassesFromContext(ctx)
		ctx, _ = templ.RenderedScriptsFromContext(ctx)
		if d.IsTrue() {
			_, err = io.WriteString(w, templ.EscapeString("True"))
			if err != nil {
				return err
			}
		} else {
			_, err = io.WriteString(w, templ.EscapeString("False"))
			if err != nil {
				return err
			}
		}
		return err
	})
}

