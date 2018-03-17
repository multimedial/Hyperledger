var username="chris";
var password = "secret";

baseURL = "http://localhost:3000/api/fabric/1_0/"

function invokeChainMethod(method, url, dataObj, callbackFunction) {

        $.ajax({
            type: method,
            url: baseURL + url,
            dataType: 'json',
            beforeSend: function (xhr) {
                xhr.setRequestHeader("Authorization", "Basic " + btoa(username + ":" + password));
            },
            data: dataObj,
            success: function (_data) {
                // isolate the transaction id from it
                setTimeout(function () { loadTransaction(_data["transactionID"], callbackFunction)}, 3500);
            },
            error: function (_data) {
                data = eval(_data.transactionEnvelope.payload.data.actions["0"].payload.action.proposal_response_payload.extension.response.payload);
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
            assignFunction(eval(data.transactionEnvelope.payload.data.actions["0"].payload.action.proposal_response_payload.extension.response.payload))
        },
        error: function (data) {
            console.log("ERROR\n"+ JSON.stringify(data));
        }
    });
}