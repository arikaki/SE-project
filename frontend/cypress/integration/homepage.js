
describe("renders the home page", () => {
  it("renders correctly", () => {
    cy.visit("/");
    cy.findAllByText("HOME").should("exist");
    cy.findAllByText("NOTIFICATIONS").should("exist");
    cy.findAllByText("CATEGORIES").should("exist");
    });


  // it("allows the date picker to be used", () => {
  //   cy.get("#date").type("2021-02-17");
  //   cy.findAllByText(CELTICS).should("exist");
  // });

  // it("routes to a team's page", () => {
  //   cy.get("#date").type("2021-02-17");
  //   cy.findAllByText(CELTICS).should("exist");
  //   cy.findAllByText(CELTICS).click();
  //   cy.url().should("include", "teams/2");
  //   cy.findByText("Conference: East").should("exist");
  });
