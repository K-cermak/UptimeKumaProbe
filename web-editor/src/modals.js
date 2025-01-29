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
                    data.splice(modalToDelete, 1);
                    modalToDelete = null;
                    renderTable();
                    closeModal("selector");
                }
            }
        }
    }
}