package visualize

const releasesTemplate string = `
<div id="releasesLoadingSpinner">
 <div class="row">
  <div class="col-md-6 offset-md-3" style="text-align: center;">
   <i class="fa fa-circle-o-notch fa-spin" style="font-size: 7rem"></i>
  </div>
 </div>
</div>

<table id="releasesTable" class="events-table table table-striped table-hover table-bordered" width="100%">
</table>
`

const releasesJS string = `
<script type="text/javascript">

function formatDetailsRow ( d ) {
    if (d["deployments"].length == 0) {
        return "";
    }

    var depList = "";
    //d["deployments"].forEach(function(dep){
    //  result = result + '<button class="btn btn-primary" type="button" style="margin-right: 0.5rem;">'+dep+'</button>';
    // });

    d["deployments"].forEach(function(dep){
      depList = depList + '<a href="#" class="list-group-item list-group-item-action">'+dep+'</a>';
    });

    var result = '<div class="list-group">' +
     '<a href="#" class="list-group-item list-group-item-action active">Used in Deployments:</a>' +
     depList +
     '</div>';
    return result;
};


$(function() {
    var releasesReq = {"command":"releases"}
    $.post("/api/command", JSON.stringify(releasesReq))
    .done(function(data) {
        var releases = data.Tables[0].Rows
        releases.forEach(function(release) {
          release["deployments"] = [];
          if (release["version"].endsWith("*")) {
            release["version"] = release["version"].slice(0, -1);
          }
        });

        var depsReq = {"command":"deployments"}
        $.post("/api/command", JSON.stringify(depsReq))
        .done(function(depsData) {
            var deployments = depsData.Tables[0].Rows
            deployments.forEach(function(deployment) {
              var depReleases = deployment["release_s"];
              var releasesList = depReleases.split("\n");

              releasesList.forEach(function(rv) {
                var releaseNameVersion = rv.split("/");

                releases.forEach(function(release) {
                  if (release["name"] === releaseNameVersion[0] &&
                       release["version"] === releaseNameVersion[1]) {
                     release["deployments"].push(deployment["name"]);
                  }
                });
              });
            });

            releases.forEach(function(release) {
              release["deployments_count"] = release["deployments"].length;
            });

            $("#releasesLoadingSpinner").remove();
            var table = $('#releasesTable').DataTable( {
                data: releases,
                columns: [
                    {
                        "class":          "details-control",
                        "orderable":      false,
                        "data":           null,
                        "defaultContent": ""
                    },
                    { data: "name", title: "name" },
                    { data: "version", title: "Version" },
                    { data: "commit_hash", title: "Commit Hash" },
                    { data: "deployments_count", title: "Deployments using it" }
                ],
                lengthChange: false,
                pageLength: 200,
                ordering: true,
                buttons: [ 'pageLength']
            } );

            var detailRows = [];

            $('#releasesTable tbody').on( 'click', 'tr td.details-control', function () {
                var tr = $(this).closest('tr');
                var row = table.row( tr );
                var idx = $.inArray( tr.attr('id'), detailRows );

                if ( row.child.isShown() ) {
                    tr.removeClass( 'details' );
                    row.child.hide();

                    // Remove from the 'open' array
                    detailRows.splice( idx, 1 );
                }
                else {
                    tr.addClass( 'details' );
                    row.child( formatDetailsRow( row.data() ) ).show();

                    // Add to the 'open' array
                    if ( idx === -1 ) {
                        detailRows.push( tr.attr('id') );
                    }
                }
            } );

            // On each draw, loop over the detailRows array and show any child rows
            table.on( 'draw', function () {
                $.each( detailRows, function ( i, id ) {
                    $('#'+id+' table.details-control').trigger( 'click' );
                } );
            } );
            });
    });
});

</script>
`
