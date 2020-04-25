const tbody = document.querySelector('#tableBody')
const queryString = new URLSearchParams(window.location.search);

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
