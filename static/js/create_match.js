document.getElementById("tournament").addEventListener("change", async function() {
  const tournamentId = this.value;
  const teamASelect = document.getElementById("teamA");
  const teamBSelect = document.getElementById("teamB");

  if (!tournamentId) {
    teamASelect.innerHTML = '<option>Seleccione un torneo primero</option>';
    teamBSelect.innerHTML = '<option>Seleccione un torneo primero</option>';
    teamASelect.disabled = true;
    teamBSelect.disabled = true;
    return;
  }

  const res = await fetch(`/tournaments/${tournamentId}/teams`);
  const teams = await res.json();

  if (!Array.isArray(teams) || teams.length === 0) {
      teamASelect.innerHTML = '<option>No hay equipos disponibles</option>';
      teamBSelect.innerHTML = '<option>No hay equipos disponibles</option>';
      teamASelect.disabled = true;
      teamBSelect.disabled = true;
      return;
    }

  const options = teams.map(t => `<option value="${t.ID}">${t.Name}</option>`).join("");
  teamASelect.innerHTML = options;
  teamBSelect.innerHTML = options;
  teamASelect.disabled = false;
  teamBSelect.disabled = false;
});
