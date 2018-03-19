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
                console.log(_data)
                // isolate the transaction id from it
                setTimeout(function () { loadTransaction(_data["transactionID"], callbackFunction)}, 3500);
            },
            error: function (_data) {
                alert("ERROR\n"+JSON.stringify(_data));
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
        success: function (_data) {

            payload = _data.transactionEnvelope.payload.data.actions[0].payload.action.proposal_response_payload
            event = payload.extension.events
            data = payload.extension.response.payload

            if (event.event_name!='')
            {
                $('#documents')[0].contentWindow.postMessage(event, "*")
                return
            }

            try {
                data = eval(data)
            }
            catch(Exception)
            {}

            assignFunction(data)

        },
        error: function (data) {
            console.log("ERROR\n"+ JSON.stringify(data));
        }
    });
}

function bin2string(array){
    var result = "";
    for(var i = 0; i < array.length; ++i){
        result+= (String.fromCharCode(array[i]));
    }
    return result;
}