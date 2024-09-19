const paths = require("./path.json");

module.exports = {
  extends: "lighthouse:default",
  settings: {
    onlyAudits: ["first-meaningful-paint", "speed-index", "interactive"],
  },
  ci: {
    collect: {
      url: paths.map((path) => `http://localhost:4173${path}`),
      startServerCommand: "cd frontend && npm run start",
    },
    upload: {
      target: "temporary-public-storage",
    },
  },
};
