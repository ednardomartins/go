var app = angular.module("GerenciadorFinanceiroTransporteCarga", []);
 
app.controller("CustoTransporteCargaController", function($scope, $http) {
	$scope.valor = "";
	$scope.veiculos = [];
	$scope.erros = [];
    $scope.detalheTransporteForm = {
    		distanciaRodoviaPavimentada: "",
    		distanciaRodoviaNaoPavimentada: "",
    		veiculoUtilizado: "",
    		cargaTransportada: ""
    };
    $scope.labelErros ="";
    
    _loadComboVeiculos()
    
    $scope.submitDetalheTransporte = function() {
        var method = "";
        var url = "";
        method = "POST";
        url = 'http://localhost:8080/v1/calcularTransporte';
        $http({
            method: method,
            url: url,
            data: angular.toJson($scope.detalheTransporteForm),
            headers: {
                'Content-Type': 'application/json'
            }
        }).then(_success, _error);
    };
    
    function _loadComboVeiculos() {
        $http({
            method: 'GET',
            url: 'http://localhost:8080/v1/veiculos'
        }).then(
            function(res) { // success
                $scope.veiculos = res.data;
                console.log(res.data);
            },
            function(res) { // error
                console.log("Error: " + res.status + " : " + res.data);
            }
        );
    }
  
    function _success(res) {
    	 $scope.erros = [];
    	 $scope.labelErros ="";
         $scope.valor = "Custo Total viagem: " + res.data.toLocaleString('pt-br',{style: 'currency', currency: 'BRL'});
    }
 
    function _error(res) {
    	$scope.labelErros ="Erros";
        $scope.erros = res.data;
    }
 
});
