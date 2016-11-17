/*
	Package processout offers bindings for the ProcessOut API —
	https://www.processout.com
	The full documentation of the API, along with Go examples, can be found on
	our website: https://docs.processout.com
	It is recommended to use the versionned version of the package (with the import
	path that starts with gopkg.in) instead of the GitHub repository.
	To get started, you just need your API credentials, that you can find in your
	project's settings.  Here is a simple example that creates an invoice and prints
	its URL:
	 p := processout.New("<project-id>", "<project-secret>")
	 iv, err := p.NewInvoice(&processout.Invoice{
		Name:     "Test item",
		Amount:   "10.00",
		Currency: "USD",
	 }).Create()
	 if err != nil {
		panic(err)
	 }
	 fmt.Println(iv.URL)
*/
package processout
