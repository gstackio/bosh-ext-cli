package visualize

const linksDeploymentDependenciesTemplate = `
<div id="linksConsumersDetailedLoadingSpinner">
 <div class="row">
  <div class="col-md-6 offset-md-3" style="text-align: center;">
   <i class="fa fa-circle-o-notch fa-spin" style="font-size: 7rem"></i>
  </div>
 </div>
</div>

<div id="linksConsumersDetailedContents" class="invisible">
<div class="row">
	<div class="col-md-7">
		<div class="card text-white bg-primary mb-3">
		  <div class="card-body">
			<p class="card-text">
              What deployments consume links from this deployment
            </p>
			 <div class="form-group" style="color: black;">
				<select class="form-control" id="deployment-select" data-placeholder="Choose a deployment">
				</select>
			 </div>
		  </div>
		</div>
	</div>
	<div class="col-md-5">
	</div>
</div>
 <div class="row">
	 <div class="col-md-12">
	   <div id="noChartMessage"></div>
	 </div>
 </div>
 <div id="chartHeaders" class="row invisible">
	 <div class="col-md-6">
	   <h4>Provider Deployment</h4>
	 </div>
	 <div class="col-md-6 pull-right">
	   <h4 class="pull-right">Consumer Deployments</h4>
	 </div>
 </div>
 <div class="row">
	 <div class="col-md-12">
	   <div id="chart"></div>
	 </div>
 </div>
</div>
`

const linksDeploymentDependenciesJS = `
<script src="https://cdnjs.cloudflare.com/ajax/libs/chosen/1.8.7/chosen.jquery.min.js" type="text/javascript" charset="utf-8"></script>
<script src="https://d3js.org/d3.v3.min.js"></script>
<script src="https://cdn.rawgit.com/newrelic-forks/d3-plugins-sankey/master/sankey.js"></script>
<script src="https://cdn.rawgit.com/misoproject/d3.chart/master/d3.chart.min.js"></script>
<script src="https://cdn.rawgit.com/q-m/d3.chart.sankey/master/d3.chart.sankey.min.js"></script>

<style type="text/css">
pre {outline: 1px solid #ccc; padding: 5px; margin: 5px; }
.string { color: green; }
.number { color: darkorange; }
.boolean { color: blue; }
.null { color: magenta; }
.key { color: red; }

#chart {
height: 500px;
font: 13px sans-serif;
}
.node rect {
fill-opacity: .9;
shape-rendering: crispEdges;
stroke-width: 0;
}
.node text {
text-shadow: 0 1px 0 #fff;
}
.link {
fill: none;
stroke: #000;
stroke-opacity: .2;
}

</style>

<script type="text/javascript">

function getUrlParameter(sParam) {
    var sPageURL = decodeURIComponent(window.location.search.substring(1)),
        sURLVariables = sPageURL.split('&'),
        sParameterName,
        i;

    for (i = 0; i < sURLVariables.length; i++) {
        sParameterName = sURLVariables[i].split('=');

        if (sParameterName[0] === sParam) {
            return sParameterName[1] === undefined ? true : sParameterName[1];
        }
    }
}

function populateDeploymentsDropDown (deploymentsList) {
    var deploymentListDropDown = $("#deployment-select");
    deploymentListDropDown.find('option').remove();
    deploymentListDropDown.append($("<option />"));
    $.each(deploymentsList, function() {
        deploymentListDropDown.append($("<option />").val(this.name).text(this.name));
    });
    deploymentListDropDown.chosen();
}

function drawGraph(currentDeployment, consumerDeploymentsList) {
    var colors = {
        'environment':         '#edbd00',
        'social':              '#367d85',
        'animals':             '#97ba4c',
        'health':              '#f5662b',
        'research_ingredient': '#3f3e47',
        'fallback':            '#edbd00'
    };

    var chart = d3.select("#chart").append("svg").chart("Sankey.Path");
    chart
        .name(label)
        .colorNodes(function(name, node) {
            return color(node, 1) || colors.fallback;
        })
        .colorLinks(function(link) {
            return color(link.source, 4) || color(link.target, 1) || colors.fallback;
        })
        .nodeWidth(15)
        .nodePadding(10)
        .spread(true)
        .iterations(0)
        .draw(consumerDeploymentsList);
    function label(node) {
        return node.name.replace(/\s*\(.*?\)$/, '');
    }
    function color(node, depth) {
        var id = node.id.replace(/(_score)?(_\d+)?$/, '');
        if (colors[id]) {
            return colors[id];
        } else if (depth > 0 && node.targetLinks && node.targetLinks.length == 1) {
            return color(node.targetLinks[0].source, depth-1);
        } else {
            return null;
        }
    }
}

function buildGraphInput(currentDeployment, consumerDeploymentsList) {
    var result = {
        nodes: [
            {
                name: currentDeployment,
                id: "SOURCE_DEPLOYMENT"
            }
        ],
        links: []
    };

    consumerDeploymentsList.forEach(function(element) {
        result.nodes.push({name: element, id: element});
    });

    consumerDeploymentsList.forEach(function(element, index) {
        result.links.push({source: 0, target: index+1, value: 1});
    });

    return result;
}

$(function() {
    var linkConsumersReq = {"command":"curl", "arguments": [{"name": "path", "value": "/deployments"}]};
    $.post("/api/command", JSON.stringify(linkConsumersReq))
        .done(function(deploymentsList) {
            populateDeploymentsDropDown(deploymentsList);

            var currentDeploymentName = getUrlParameter('deployment');
            if (currentDeploymentName) {
                // Get all the consumers of this deployment
                // For each deployment, get the link providers
                // when that is done, get the links of the deployment
                // match provider for each link consumer

                // Get all link consumers of this deployment
                var linkConsumersReq = {"command":"curl", "arguments": [{"name": "path", "value": "/link_providers?deployment="+currentDeploymentName}]};
                $.post("/api/command", JSON.stringify(linkConsumersReq))
                    .done(function(data) {
                        var currentDeploymentProviders = data.reduce(function(map, obj) {
                            map[obj["id"]] = obj;
                            return map;
                        }, {});

                        var numberOfTotalRequests = deploymentsList.length; // number of total requests
                        var aggregatedLinksList = {}; // key: deployment name, value: list of links

                        $(deploymentsList).each(function() {
                            var deploymentName = this.name;
                            var linksReq = {"command":"curl", "arguments": [{"name": "path", "value": "/links?deployment="+deploymentName}]};
                            $.post("/api/command", JSON.stringify(linksReq))
                                .done(function(data) {
                                    aggregatedLinksList[deploymentName] = data
                                    numberOfTotalRequests -= 1;

                                    if(numberOfTotalRequests === 0) {

                                        // all links of all deployments were retrieved
                                        // select links which have their link_provider_id belongs to the current deployment link providers
                                        var result = [];
                                        $.each( aggregatedLinksList, function( someDeploymentName, linksList ) {
                                            $(linksList).each(function (i, linkObj) {
                                                if (currentDeploymentProviders[linkObj["link_provider_id"]]) {
                                                    result.push(someDeploymentName);
                                                }
                                            })
                                        });

                                        var sortedUniqueDepList = Array.from(new Set(result)).sort();
                                        console.log(sortedUniqueDepList);
                                        console.log(buildGraphInput(currentDeploymentName, sortedUniqueDepList));
                                        $("#linksConsumersDetailedLoadingSpinner").remove();
                                        $("#linksConsumersDetailedContents").removeClass("invisible").addClass("visible");

                                        var graphInput = buildGraphInput(currentDeploymentName, sortedUniqueDepList);
                                        if (graphInput["links"].length === 0) {
                                            $("#noChartMessage").text("No links are consumed from this deployment '" + currentDeploymentName + "'");
                                        } else {
                                            $("#chartHeaders").removeClass("invisible").addClass("visible");
                                            drawGraph(currentDeploymentName, buildGraphInput(currentDeploymentName, sortedUniqueDepList));
                                        }
                                    }
                                });
                        });
                    });
            } else {
                $("#linksConsumersDetailedLoadingSpinner").remove();
                $("#linksConsumersDetailedContents").removeClass("invisible").addClass("visible");
            }
        });

    $("#deployment-select").change(function () {
        window.location.href = "/links-deployments-dependencies?deployment=" + $("#deployment-select").find(":selected").text();
    });
});

</script>
`
