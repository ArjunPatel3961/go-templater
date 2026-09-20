# Tiny {{var}} templater

```
templater.go
```
When you are wiring up OTP delivery flows or rendering transactional SMS bodies, you generally want to avoid pulling in a heavy template engine. This utility resolves {{name}} placeholders from a data object using strictly the Go standard library.

Check the test file sitting next to the implementation for concrete usage. It covers the messy edge cases we actually hit in production, like missing dictionary keys or unexpected nil values. You get a lightweight resolver without any external services or extra dependencies to manage in your build pipeline.