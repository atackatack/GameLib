function SearchAddGameHLTB() {
    delay(function () {
        let value = $('#inputNewGame').val()
        $.ajax({
            url: '/hltb/search/',
            method: 'GET',
            data: {
                name: value
            },
            dataType: 'json',
            success: function (body) {
                loadSearchGameHLTB('#addListGameHLTB', body)
            }
        });
    }, 200);
}

function switchInputGameHLTB() {
    let selectedDiv;
    $(document).on('click', '.htlbGameAdd', function (event) {
        this.classList.add("active")
        if (selectedDiv == this) {
            selectedDiv.classList.remove("active")
            $('.gameHLTB').attr("value", "0")
            $('#searchAddGameHLTB').attr("placeholder", "Игра из HLTB")
            selectedDiv = null
            return
        }

        if (selectedDiv != null) {
            selectedDiv.classList.remove("active")
        }

        selectedDiv = this
        $('.gameHLTB').attr("value", this.getAttribute("value"))
        $('#searchAddGameHLTB').attr("placeholder",
            this.getAttribute("name") + " | " + this.getAttribute("value"))
        $('#inputNewGame').val(this.getAttribute("name"))
    });
}

function loadSearchGameHLTB(elem, body) {
    if (body['data'] === null) {
        $(elem).html('');
        return
    }

    let rows = ""
    body['data'].forEach(function (obj) {
        rows += `<div href="#" class="list-group-item list-group-item-action d-flex gap-3 py-3 htlbGameAdd"
               aria-current="true" value=${obj.game_id} name="${obj.game_name}">
               <img src="${obj.game_image}" alt="twbs" width="80" height="120"
                     class="rounded flex-shrink-0">
               <div>
                    <h6><b>${obj.game_name}</b></h6>
                    <p class="mb-0 opacity-75"><b>Game ID:</b> ${obj.game_id}</p>
                    <p class="mb-0 opacity-75"><b>Main Time:</b> ${obj.comp_main} Hours</p>
                    <p class="mb-0 opacity-75"><b>Full Time:</b> ${obj.comp_plus} Hours</p>
               </div>
            </div>`
    });

    $(elem).html(rows);
}

switchInputGameHLTB();