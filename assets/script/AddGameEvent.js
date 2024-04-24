function getInfoForAdd(hltd_id) {
    let obj;

    if ($('#gridForGame').get(0).files[0]) {
        obj = {
            hltb_id: hltd_id,
            name: $("#inputNewGame").val(),
            done: $(".flexChecked").is(":checked"),
            image: $('#gridForGame').get(0).files[0].name,
            find_grid: $(".findHltbGrid").is(":checked"),
        };
    } else {
        obj = {
            hltb_id: hltd_id,
            name: $("#inputNewGame").val(),
            done: $(".flexChecked").is(":checked"),
            find_grid: $(".findHltbGrid").is(":checked"),
        };
    }

    return JSON.stringify(obj);
}

function clearInputForm() {
    $("#gameHLTB").attr('value', 0);
    $('#inputNewGame').val("");
    $('#gridForGame').val("");
    $('#searchAddGameHLTB').attr('placeholder', 'Игра из HLTB');
    $('#addListGameHLTB').html("");
    $(".flexChecked").prop('checked', false);
    $(".findHltbGrid").prop('checked', true);
}

function createGame() {
    let hltd_id = Number($('#gameHLTB').attr("value"))
    $.ajax({
        url: '/api/v1/game/',
        method: 'POST',
        dataType: 'json',
        data: getInfoForAdd(hltd_id),
        statusCode: {
            200: function () {
                alert("Игра уже есть в списке");
            },
            201: async function () {
                await saveImg($('#gridForGame').get(0).files[0]);
                clearInputForm();
                refreshTable();
                $('#addGameModal').modal('toggle');
            },
            400: function () {
                alert("Что-то пошло не так!");
            },
            422: function () {
                alert("Имя игры слишком большое или пустое");
            }
        }
    });
}

$(document).on('click', '#buttonAdd', function (event) {
    createGame()
});

$('#inputNewGame').bind("enterKey", function (e) {
    createGame()
});

$('#inputNewGame').keyup(function (e) {
    if (e.keyCode == 13) {
        $(this).trigger("enterKey");
    }
});

$('#addGameModal').on('hidden.bs.modal', function () {
    clearInputForm()
})