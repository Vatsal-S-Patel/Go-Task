let prevButton = document.getElementById("prev");
let nextButton = document.getElementById("next");
let counter = document.getElementById("counter");
let search = document.getElementById("search");

let sort = "";
let name = document.getElementById("trainname");
name.addEventListener("click", (e) => {
  sort = "name";
  loadPage();
});

let searchText = "";
let count = 1;

counter.innerText = count;

loadPage();

prevButton.addEventListener("click", (e) => {
  if (count <= 1) {
    return;
  }
  count--;
  loadPage();
  console.log(count);
});

nextButton.addEventListener("click", (e) => {
  count++;
  loadPage();
  console.log(count);
});

search.addEventListener("keypress", (e) => {
  if (e.key == "Enter") {
    searchText = e.target.value;
    count = 1;
    loadPage();
  }
  // setTimeout(() => {
  // }, 1000);
  // clearTimeout()
});

function loadPage() {
  let li = "";
  fetch(
    "http://localhost:8080/api/trains/?page=" +
      count +
      "&search=" +
      searchText +
      "&sort" +
      sort
  )
    .then((res) => res.json())
    .then((trains) => {
      if (trains == null && count != 1) {
        --count;
        return;
      }
      counter.innerText = count;
      trains.forEach((train) => {
        li += `<tr>
                <td>${train.sno} </td> 
                <td>${train.trainNo}</td>         
                <td>${train.name}</td>    
                <td>${train.source}</td>    
                <td>${train.destination}</td>    
        </tr>`;
      });
      document.getElementById("trains-list").innerHTML = li;
    })
    .catch((err) => {
      console.log(err);
    });
}
