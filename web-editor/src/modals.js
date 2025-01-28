var resetFavModal = {
    global : {
        closable : true,
        size : "md",
        scrollable : false,
        position : "center",
    },
    header : {
        title : "Clear scan",
        closeButton: true,
    },
    main : {
        content : "Are you sure you want to delete this scan?",
    },
    footer : {
        buttons : {
            close : {
                text : "Cancel",
                type : "secondary",
                function : "close",
            },
            function : {
                text : "<i class='bi bi-trash3' ms-0 me-1'></i> Delete",
                type : "danger",
                function : "function",
                dataset : function() {
                    console.log(modal_to_detele);
                    data.splice(modal_to_detele, 1);
                    modal_to_detele = null;
                    renderTable();
                    closeModal("selector");
                }
            }
        }
    }
}