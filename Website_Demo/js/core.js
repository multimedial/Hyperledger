var username="chris";
var password = "secret";

baseURL = "http://localhost:3000/api/fabric/1_0/"

function invokeChainMethod(method, url, dataObj, callbackFunction) {


        $.ajax({
            type: method,
            url: baseURL + url,
            dataType: 'json',
            beforeSend: function (xhr) {
                xhr.setRequestHeader("Authorization", "Basic " + btoa("chris:secret"));
            },
            data: dataObj,
            success: function (data) {
                // isolate the transaction id from it
                setTimeout(function () { loadTransaction(data["transactionID"], callbackFunction)}, 3000);
            },
            error: function (data) {
                alert("ERROR\n"+JSON.stringify(data));
            }
    });

    }

function loadTransaction(transactionid, assignFunction) {

    url = "http://localhost:3000/api/fabric/1_0/channels/vertraulich/transactions/"+transactionid

    $.ajax({
        type: "GET",
        url: url,
        dataType: 'json',
        beforeSend: function (xhr) {
            xhr.setRequestHeader("Authorization", "Basic " + btoa("chris:secret"));
        },
        success: function (data) {
            console.log(data);
            _data = eval(data.transactionEnvelope.payload.data.actions["0"].payload.action.proposal_response_payload.extension.response.payload);
            assignFunction(_data)

        },
        error: function (data) {
            console.log("ERROR\n"+ JSON.stringify(data));
        }
    });
}