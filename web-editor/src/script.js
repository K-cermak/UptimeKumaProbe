setTimeout(function () {
    var popoverTriggerList = [].slice.call(document.querySelectorAll('[data-bs-toggle="popover"]'))
    popoverTriggerList.map(function (popoverTriggerEl) {
        return new bootstrap.Popover(popoverTriggerEl)
    })
}, 200);

//data
const table = document.querySelector("#config tbody");
var modal_to_detele = null;

var data = [
    {
      name: "service_name1",
      type: "ping",
      address: "10.0.0.4",
      timeout: 50,
    },
    {
      name: "service_name2",
      type: "http",
      address: "address",
      timeout: 50,
    },
    {
      name: "service_name3",
      type: "http",
      address: "address",
      timeout: 50,
      keyword: "Some words",
    },
    {
      name: "service_name4",
      type: "http",
      address: "address",
      timeout: 50,
      status_code: "200;201;202",
    },
    {
      name: "service_name5",
      type: "http",
      address: "address",
      timeout: 50,
      status_code: "200;201",
      keyword: "Some words",
    },
];

function renderTable() {
    table.innerHTML = "";

    for (let i = 0; i < data.length; i++) {
        let row = document.createElement("tr");

        let index = document.createElement("td");
        index.innerHTML = "#" + (i + 1);
        row.appendChild(index);

        let name = document.createElement("td");
        let name_input = document.createElement("input");
        name_input.classList.add("form-control");
        name_input.value = data[i].name;
        name.appendChild(name_input);
        row.appendChild(name);

        let type = document.createElement("td");
        let select = document.createElement("select");

        let options = ["ping", "http"];
        for (let j = 0; j < options.length; j++) {
            let option = document.createElement("option");
            option.value = options[j];
            option.text = options[j];
            select.appendChild(option);
        }
        select.value = data[i].type;
        select.classList.add("form-select");
        type.appendChild(select);
        row.appendChild(type);

        let address = document.createElement("td");
        let address_input = document.createElement("input");
        address_input.classList.add("form-control");
        address_input.value = data[i].address;
        address.appendChild(address_input);
        row.appendChild(address);

        let timeout = document.createElement("td");
        let timeout_input = document.createElement("input");
        timeout_input.classList.add("form-control");
        timeout_input.type = "number";
        timeout_input.value = data[i].timeout;
        timeout.appendChild(timeout_input);
        row.appendChild(timeout);

        let status = document.createElement("td");
        let status_input = document.createElement("input");
        status_input.classList.add("form-control");
        if (data[i].type === "http") {
            status_input.value = data[i].status_code || "";
        } else {
            status_input.value = "-";
            status_input.disabled = true;
        }
        status.appendChild(status_input);
        row.appendChild(status);

        let keyword = document.createElement("td");
        let keyword_input = document.createElement("input");
        keyword_input.classList.add("form-control");
        if (data[i].type === "http") {
            keyword_input.value = data[i].keyword || "";
        } else {
            keyword_input.value = "-";
            keyword_input.disabled = true;
        }
        keyword.appendChild(keyword_input);
        row.appendChild(keyword);

        let detele = document.createElement("td");
        let button = document.createElement("button");
        button.classList.add("btn", "btn-danger");
        button.innerHTML = "<i class='bi bi-x-circle'></i>";
        button.onclick = function () {
            modal_to_detele = i;
            genModal(resetFavModal);
        };
        detele.appendChild(button);
        row.appendChild(detele);
        
        table.appendChild(row);
    }
}


//RUNS
renderTable();