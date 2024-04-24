function refreshTable() {
    updateGameStatus()
    $.getJSON('/api/v1/game/all', function (body) {
        let rows = "";

        if (body['data'] == null) {
            $('#gameTable').html(``);
            return;
        }

        body['data'].forEach(function (obj) {
            rows += `<div class="card" data-name="${obj.name}">
                        <div class="poster">
                            <input class="form-check-input statusCheckBox" type="checkbox"
                                   value=${obj.id} `

            if (obj.done) {
                rows += `checked>`;
            } else {
                rows += `>`;
            }

            if (obj.image_url) {
                rows += `<img src="${obj.image_url}" alt="${obj.name}">`
            } else {
                rows += `<img src="../assets/static/tmpGrid.png" alt="${obj.name}">`
            }
            rows += `<button type="button" id="buttonDeleteElem"
                            class="btn btn-sm text-center deleteElemButton" name="buttonDelete"
                            value="${obj.id}">
                    </button>
                    <button type="button" id="buttonUpdateElem" data-bs-toggle="modal" data-bs-target="#updateGameModal"
                            class="btn btn-sm text-center updateElemButton" name="buttonUpdate"
                            value="${obj.id}">
                    </button>`

            if (obj.favorite) {
                rows += `<button class="btn btn-sm text-center favoriteElemButton favoriteFill" id="favoriteElemButton"
                        value="${obj.id}"></button>`
            } else {
                rows += `<button class="btn btn-sm text-center favoriteElemButton" id="favoriteElemButton"
                        value="${obj.id}"></button>`
            }

            rows += `</div>
                    <div class="details" id="gameDetails">
                        <h1>${obj.name}</h1>
                        <div class="row">
                            <div class="col-md-6 mx-auto" style="font-size: .75rem">Main: ${obj.hltb_main_time} H</div>
                            <div class="col-md-6 mx-auto" style="font-size: .75rem">Full: ${obj.hltb_full_time} H</div>
                        </div>
                    </div>
                </div>`;
        });

        $('#gameTable').html(rows);
        $.each($("#gameTable div.card"), function () {
            if (this.getAttribute("data-name").toLowerCase().indexOf($("#search").val().toLowerCase()) == -1) {
                $(this).hide();
            } else {
                $(this).show();
            }
        });
    });
}

function updateGameStatus() {
    $('#doneGameCounter').html(getDoneGameCount())
    $('#allGameCounter').html(getAllGameCount())
}