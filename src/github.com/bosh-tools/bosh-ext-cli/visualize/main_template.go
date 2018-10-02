package visualize

func generateTemplate(body, customJS string) string {

	return `
<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
  <title>{{.Title}} Template</title>
  <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css" integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO" crossorigin="anonymous">
  <link href="https://stackpath.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css" rel="stylesheet" integrity="sha384-wvfXpqpZZVQGK6TAh5PVlGOfQNHSoD2xbE+QkPxCAFlNEevoEH3Sl0sibVcOQVnN" crossorigin="anonymous">
  <link rel="stylesheet" type="text/css" href="https://cdn.datatables.net/1.10.16/css/dataTables.bootstrap4.min.css"/>
  <link rel="stylesheet" type="text/css" href="https://cdn.datatables.net/buttons/1.5.1/css/buttons.bootstrap4.min.css"/>
  <link rel="stylesheet" type="text/css" href="https://cdnjs.cloudflare.com/ajax/libs/chosen/1.8.7/chosen.min.css"/>
  <link href="css/sb-admin.css" rel="stylesheet">
</head>

<body>

  <!-- Navigation-->` +
		navTemplate +
		`<!-- END Navigation-->

    <main role="main" class="container" style="max-width: 2000px; padding-top: 60px; margin-right: 15px; margin-left: 15px;">  {{.Title}}` +
		body +
		`<!-- ============================================= -->
    </main>


    <!-- ============================================= -->
    <!-- ============================================= -->
    <!-- ============================================= -->
    <!-- ============================================= -->
    <!-- Bootstrap core JavaScript-->
    <script 
      src="https://code.jquery.com/jquery-3.3.1.min.js"
      integrity="sha256-FgpCb/KJQlLNfOu91ta32o/NMZxltwRo8QtmkMRdAu8="
      crossorigin="anonymous"></script>
    <script 
      src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.12.9/umd/popper.min.js" 
      integrity="sha384-ApNbgh9B+Y1QKtv3Rn7W3mgPxhU9K/ScQsAP7hUibX39j7fakFPskvXusvfa0b4Q" 
      crossorigin="anonymous"></script>
    <script 
      src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/js/bootstrap.min.js" 
      integrity="sha384-JZR6Spejh4U02d8jOt6vLEHfe/JQGiRRSQQxSfFWpi1MquVdAyjUar5+76PVCmYl" 
      crossorigin="anonymous"></script>
    <script 
      src="https://cdnjs.cloudflare.com/ajax/libs/jquery-easing/1.4.1/jquery.easing.min.js" 
      integrity="sha256-H3cjtrm/ztDeuhCN9I4yh4iN2Ybx/y1RM7rMmAesA0k=" 
      crossorigin="anonymous"></script>
    <script 
      src="https://cdnjs.cloudflare.com/ajax/libs/underscore.js/1.8.3/underscore-min.js" 
      integrity="sha256-obZACiHd7gkOk9iIL/pimWMTJ4W/pBsKu+oZnSeBIek=" 
      crossorigin="anonymous"></script>  

    <script 
      type="text/javascript" 
      src="https://cdn.datatables.net/1.10.16/js/jquery.dataTables.min.js">
    </script>

    <script 
      type="text/javascript" 
      src="https://cdn.datatables.net/1.10.16/js/dataTables.bootstrap4.min.js">
    </script>

    <script 
      type="text/javascript" 
      src="https://cdn.datatables.net/buttons/1.5.1/js/dataTables.buttons.min.js">
    </script>

    <script 
      type="text/javascript" 
      src="https://cdn.datatables.net/buttons/1.5.1/js/buttons.bootstrap4.min.js">
    </script>

    <script 
      type="text/javascript" 
      src="https://cdn.datatables.net/buttons/1.5.1/js/buttons.colVis.min.js">
    </script>


    <!-- Custom scripts for all pages-->
    <script src="js/sb-admin.min.js"></script>` +
		customJS + `
    <script type="text/javascript">
      envReq = {"command":"env"}
      $.post("/api/command", JSON.stringify(envReq))
        .done(function(data) {
          $("#bosh-env-btn").text(data["Lines"][0]);
        });
    </script>
  </div>
</body>

</html>

`
}
