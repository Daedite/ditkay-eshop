package wire

import (
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/config"
	factory3 "github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/models/gorm/factory"
	gorm2 "github.com/ESPOIR-DITE/ditkay-eshop/pkg/db/repository/gorm"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity/factory"
	controller2 "github.com/ESPOIR-DITE/ditkay-eshop/pkg/http/controller"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/http/server"
	impl2 "github.com/ESPOIR-DITE/ditkay-eshop/pkg/service/media/impl"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/service/media_type/impl"
	product "github.com/ESPOIR-DITE/ditkay-eshop/pkg/service/product/impl"
	impl3 "github.com/ESPOIR-DITE/ditkay-eshop/pkg/service/product_media/impl"
	"gorm.io/gorm"
)

func wire(config config.ServiceConfiguration, gorm *gorm.DB) server.HttpServer {

	mediaFactory := factory3.NewMediaFactory()
	productFactory := factory3.NewProductFactory()
	productMediaFactory := factory3.NewProductMediaFactory()
	mediaTypeFactory := factory3.NewMediaTypeFactory()

	productRepository := gorm2.NewProductRepositoryImpl(gorm, productFactory)
	mediaRepository := gorm2.NewMediaRepositoryImpl(gorm, mediaFactory)
	productMediaRepository := gorm2.NewProductMediaRepositoryImpl(gorm, productMediaFactory)
	mediaTypeRepository := gorm2.NewMediaTypeRepositoryImpl(gorm, mediaTypeFactory)

	productService := product.NewProductServiceImpl(productRepository)
	mediaTypeService := impl.NewMediaTypeServiceImpl(mediaTypeRepository)
	mediaService := impl2.NewMediaService(mediaRepository)
	productMediaService := impl3.NewProductMediaService(productMediaRepository)

	mediaEntityFactory := factory.NewMediaFactoryImpl()
	mediaTypeEntityFactory := factory.NewMediaTypeFactoryImpl()
	productEntityFactory := factory.NewProductFactoryImpl()
	productMediaEntityFactory := factory.NewProductMediaFactoryImpl()

	controller := controller2.NewDitKayEshopApiController(
		mediaService,
		mediaTypeService,
		productService,
		productMediaService,
		mediaEntityFactory,
		mediaTypeEntityFactory,
		productEntityFactory,
		productMediaEntityFactory,
	)
	return server.NewHttpServerImpl(config.AppConfig(), controller)

}
