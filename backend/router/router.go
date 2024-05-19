package router

import (
	"mithuorganics/controller"

	"github.com/gin-gonic/gin"
)

// Router creates and configures the Gin router.
func Router() *gin.Engine {
	router := gin.Default()
	router.Use(controller.CorsMiddleware())
	//Serve static files for specific routes
	router.Static("/admin/signin", "./frontend/admin/signin")
	router.Static("/admin/dashboard","./frontend/admin/dashboard")
	router.Static("/home", "./frontend/home")
	router.Static("/signup", "./frontend/signup")
	router.Static("/signin", "./frontend/signin")
	router.Static("/seller/signin","./frontend/seller/signin")
	router.Static("/seller/signup", "./frontend/seller/signup")
	router.Static("/seller/dashboard", "./frontend/seller/dashboard")

	// Define your routes
	router.GET("/getallcustomerdata", controller.GetallCustomerdata)
	router.GET("/getallinventorydata", controller.Getinventorydata)
	router.POST("/getallsellerdata", controller.Getallsellerdata)
	router.POST("/createseller", controller.CreateSeller)
	router.POST("/getdata", controller.GetData)
	router.POST("/create", controller.CreateProfile)
	router.POST("/deletedata", controller.Delete)
	router.POST("/addtocart", controller.Addtocart)
	router.POST("/verifyemail", controller.VerifyEmail)
	router.POST("/login", controller.Login)
	router.POST("/products", controller.Products)
	router.POST("/updatecart", controller.UpdateCart)
	router.POST("/inventory", controller.Inventory)
	router.POST("/search", controller.Search)
	router.POST("/update", controller.Update)
	router.POST("/sellercheck", controller.CheckSeller)
	router.POST("/deleteproduct", controller.DeleteProduct)
	router.POST("/deleteproductbyseller", controller.DeleteProductBySeller)
	router.POST("/updateproductbyseller", controller.UpdateProductBySeller)
	router.POST("/insertsellerfeedback", controller.InsertSellerFeedback)
	router.POST("/insertcustomerfeedback", controller.InsertCustomerFeedback)
	router.POST("/getfeedback", controller.GetFeedback)
	router.GET("/sellerfeedback", controller.SellerFeedback)
	router.GET("/customerfeedback", controller.CustomerFeedback)
	router.POST("/deletefeedback", controller.Deletefeedback)
	router.POST("/buynow", controller.BuyNow)
	router.POST("/totalamount", controller.TotalAmount)
	router.POST("/orders",controller.Orders)
	router.POST("/deleteorder",controller.DeleteOrder)
	router.POST("/customerorders",controller.CustomerOrder)
	router.POST("/validatetoken", controller.ValidateToken)
	router.POST("/adminlogin", controller.AdminLogin)
	router.POST("/adminpage", controller.GetAllDetailsForAdmin)
	router.POST("/getworkers", controller.GetWorkers)
	router.POST("/createworker",controller.CreateWorker)
	router.POST("/createadmin",controller.CreateAdmin)
	router.POST("/addevent",controller.AddEvent)
	router.POST("/getevent",controller.GetEvent)
	router.POST("/block",controller.Block)
	router.POST("/getuseraddress",controller.GetUserAddress)
	router.POST("/getinventorydata",controller.GetInventoryData)
	router.POST("/adddeliveryaddress",controller.AddDeliveryAddress)
	router.POST("/getcustomerorder",controller.GetCustomerOrder)
	router.POST("/getcustomerorderforadmin",controller.GetCustromerOrderforAdmin)
	router.POST("/registerseller",controller.RegisterSeller)
	router.POST("/verifyselleremail",controller.VerifySellerEmail)
	router.POST("/sellerdrashboarddetails",controller.SellerDrashbordDetails)
	router.POST("/getallproducts",controller.GetAllProducts)
	router.POST("/buyedcustomer",controller.BuyedCustomer)
	router.POST("/completedorders",controller.CompletedOrders)
	router.POST("/pendingorders",controller.PendingOrders)
	router.POST("/yettodeliverorders",controller.YettoDeliverOrders)
	router.POST("/updateordertracking",controller.UpdateOrderTracking)
	router.POST("/getproductforseller",controller.GetProductData)
	router.POST("/shutdown",controller.ShutDown)
	router.POST("/cleardb",controller.ClearDB)
	router.POST("/getallnotapprovedseller",controller.GetAllNotApprovedSeller)
	router.POST("/approveseller",controller.ApproveSeller)
	router.POST("/getallorders",controller.GetAllOrders)


	return router
}