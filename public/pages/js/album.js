const deleteBtn = document.getElementById(`deleteBtn`);

deleteBtn.addEventListener('click',() => {
    const id = deleteBtn.dataset.albumId;
    console.log(id);
    deleteAlbum(id);
});

async function deleteAlbum(id) {
    const res = await fetch(`/albums/${id}/delete`,{
        method : "DELETE"
    })

    const data = await res.json();
    console.log(data);

    if(data.success) {
        result = confirm("Are you sure? If you delete this album, you cannot recover it");
        if(result) {
            document.getElementById(`albums${id}`).remove();
            window.location.href = "/albums";
        }
    } else {
        alert(data.message);
    }
}