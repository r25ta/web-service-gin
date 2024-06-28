module example.com/web-service-gin/repository

go 1.22.2

replace example.com/web-service-gin/constant => ../constant

require (
	example.com/web-service-gin/constant v0.0.0-00010101000000-000000000000
	example.com/web-service-gin/model v0.0.0-00010101000000-000000000000
	github.com/lib/pq v1.10.9
)

replace example.com/web-service-gin/model => ../model
