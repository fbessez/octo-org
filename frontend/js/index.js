const queryString = new URLSearchParams(window.location.search);
const tbody = document.querySelector('#tableBody')

function renderList(jsonData) {
  return `${jsonData.map(
    (row, index) => `<tr>
      <td class="rank"> ${index + 1}.</td>
      <td class="user"> ${row['github_username']}</td>
      <td class="metric"> ${row['total_commits']}</td>
      <td class="metric"> ${row['total_additions']}</td>
      <td class="metric"> ${row['total_deletions']}</td>
      <td class="metric"> ${(row['total_additions'] / row['total_deletions']).toFixed(2) }</td>
    `).join('')}`;
}

async function fillTableData() {
  let sortOption = queryString.get('sort');
  let queryParam = sortOption ? '?sort=' + sortOption : '';
  const URL = 'http://localhost:8090/stats' + queryParam;

  try {
    const fetchResult = fetch(new Request(URL, { method: 'GET', cache: 'reload' }));
    const response = await fetchResult;
    if (response.ok) {
      const jsonData = await response.json();
      tbody.innerHTML = renderList(jsonData);
    } else {
      tbody.innerHTML = ``;
      window.alert("API failed...go home.")
    }
  } catch (e) {
    tbody.innerHTML = ``;
    window.alert("API failed...go home.")
  }
}

fillTableData()

function handleFilterClick(e) {
  // TODO: rather than rerendering on click due to the href, just resort the fetched (once) data.
  // function sortBy(prop){
  //    return function(a,b){
  //       if (a[prop] > b[prop]){
  //           return 1;
  //       } else if(a[prop] < b[prop]){
  //           return -1;
  //       }
  //       return 0;
  //    }
  // }
  e.preventDefault();
  filterButtons.forEach(node => {
    node.classList.remove('active')
  })

  e.currentTarget.classList.add('active')
}

let filterButtons = Array.from(document.querySelectorAll('.filter-button'))
filterButtons.forEach(btn => {
  btn.addEventListener('click', handleFilterClick)
})



