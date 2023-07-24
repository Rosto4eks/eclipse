const deleteBtn = document.getElementById(`deleteBtn`);

deleteBtn.addEventListener('click',() => {
    const id = deleteBtn.dataset.albumId;
    console.log(id);
    deleteAlbum(id);
});

async function deleteAlbum(id) {
    result = confirm("Are you sure? If you delete this album, you cannot recover it");
    if(!result){
        return;
    }
    const res = await fetch(`/albums/${id}/delete`,{
        method : "DELETE"
    })

    const data = await res.json();
    console.log(data);

    if(data.success) {
        document.getElementById(`albums${id}`).remove();
        window.location.href = "/albums";
    } else {
        alert(data.message);
    }
}
