function displaySelectedImage(event, elementId) {
    const selectedImage = document.getElementById(elementId);
    const fileInput = event.target;

    if (fileInput.files && fileInput.files[0]) {
        const reader = new FileReader();
        reader.onload = function (e) {
            selectedImage.src = e.target.result;
        };

        reader.readAsDataURL(fileInput.files[0]);
    }
}

$(document).on('click', '#buttonUpdateElem', function (event) {
    let id = $(this).val()
    $.getJSON('/api/v1/game/' + id, function (body) {
        let rows = `
            <div class="col mb-2">
              <label for="updateGameNameInput" class="form-label">Название</label>
              <input type="text" class="form-control" name="game" id="updateGameNameInput" 
                        autocomplete="off" value="${body.data.name}" placeholder="Введите название игры...">
            </div>
            <div class="container">
                <div class="row">
                    <div class="col form-check mb-2">
                        <input type="checkbox" class="form-check-input" id="updateGameStatusInput"`;

        if (body.data.done) {
            rows += `checked><label class="form-check-label" for="updateGameStatusInput">Статус</label></div>`;
        } else {
            rows += `><label class="form-check-label" for="updateGameStatusInput">Статус</label></div>`;
        }

        rows += `  
                    <div class="col form-check mb-2">
                        <input type="checkbox" class="form-check-input" id="updateGridHltbInput"> <label class="form-grid-hltb-label" for="updateGridHltbInput">Поиск обложки</label>
                    </div>
                </div>
            </div>
            
            <div class="mb-2 gameHLTB" id="gameHLTB" value=0>
                <div class="input-group has-validation">
                    <span class="input-group-text">HLTB</span>
                    <input type="text" class="form-control" id="searchUpdateGameHLTB" placeholder="Название игры" required=""
                           autocomplete="off" onkeyup="updateGameListHLTB($('#searchUpdateGameHLTB').val())"/>
                </div>
            </div>
            
            <div class="mb-2 flex-column flex-md-row align-items-center justify-content-center">
                <div id="updateListGameHLTB" class="list-group overflow-auto shadow align-items-center" style="max-height: 160px;">
                </div>
            </div>
            
            <div class="inputGrid mb-2">`

        if (body.data.image_url) {
            rows += `<img id="selectedImage" src="${body.data.image_url}" class="mx-auto d-block"></div>`
        } else {
            rows += `<img id="selectedImage" src="../assets/static/tmpGrid.png" class="mx-auto d-block"></div>`
        }

        rows += `<label class="col-12 btn btn-primary btn-rounded form-label text-white m-1" for="updateGridButton">Обложка</label>
                <input type="file" class="form-control d-none" accept="image/png, image/jpeg" id="updateGridButton" onchange="displaySelectedImage(event, 'selectedImage')"/>`;

        $('#updateGameBody').html(rows);
        $('#buttonUpdate').attr('value', id);
    });
});

function getInfoForUpdate() {
    let obj;

    if ($('#updateGridButton').get(0).files[0]) {
        obj = {
            name: $("#updateGameNameInput").val(),
            hltb_id: Number($('#gameHLTB').attr("value")),
            done: $("#updateGameStatusInput").is(":checked"),
            find_grid: $("#updateGridHltbInput").is(":checked"),
            image: $('#updateGridButton').get(0).files[0].name,
        };
    } else {
        obj = {
            name: $("#updateGameNameInput").val(),
            hltb_id: Number($('#gameHLTB').attr("value")),
            done: $("#updateGameStatusInput").is(":checked"),
            find_grid: $("#updateGridHltbInput").is(":checked"),
        };
    }

    return JSON.stringify(obj);
}

function clearUpdateForm() {
    $('#updateGameNameInput').val("");
    $('#gameHLTB').attr("value", 0);
    $('#searchAddGameHLTB').attr('placeholder', 'Игра из HLTB');
    $(".updateGameStatusInput").prop('checked', false);
    $(".updateGridHltbInput").prop('checked', false);
}

function updateGame(id) {
    $.ajax({
        url: '/api/v1/game/' + id,
        method: 'PUT',
        dataType: 'json',
        data: getInfoForUpdate(),
        statusCode: {
            200: async function () {
                await saveImg($('#updateGridButton').get(0).files[0]);
                $('#updateGameModal').modal('toggle');
                refreshTable();
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

$(document).on('click', '#buttonUpdate', function (event) {
    updateGame($(this).val())
});

$('#updateGameNameInput').bind("enterKey", function (e) {
    updateGame($(this).val())
});

$('#updateGameNameInput').keyup(function (e) {
    if (e.keyCode == 13) {
        $(this).trigger("enterKey");
    }
});

$('#updateGameModal').on('hidden.bs.modal', function () {
    clearUpdateForm()
})

function updateGameListHLTB(name) {
    delay(function () {
        $.ajax({
            url: '/hltb/search/',
            method: 'GET',
            data: {
                name: name
            },
            dataType: 'json',
            success: function (body) {
                loadSearchGameHLTB('#updateListGameHLTB', body)
            }
        });
    }, 200);
}