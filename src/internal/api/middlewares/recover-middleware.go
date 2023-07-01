package middlewares

// func Recover(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		defer func() {
// 			if err := recover(); err != nil {
// 				ctx := r.Context()
// 				logger := zerolog.FromContext(ctx)
// 				stack := make([]byte, 8096)
// 				stack = stack[:runtime.Stack(stack, false)]
// 				logger.Bytes("stack", stack).Str("level", "fatal").Interface("error", err)
// 				hasRecovered, ok := ctx.Value(polymer.ContextKey("has_recovered")).(*bool)
// 				if ok {
// 					*hasRecovered = true
// 				}
// 				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 			}
// 		}()
// 		next.ServeHTTP(w, r)
// 	})
// }
