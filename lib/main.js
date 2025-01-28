function seePets(){
       fetch("http://localhost/8000/pets", {
            method: "GET",
            headers: {
                "Content-Type": "application/json"
            },
        })
        .then(res => res.json())
        .then(res => {
            console.log('Pets:', res);
        })
}


seePets();
