// ***********************************************
// This example commands.js shows you how to
// create various custom commands and overwrite
// existing commands.
//
// For more comprehensive examples of custom
// commands please read more here:
// https://on.cypress.io/custom-commands
// ***********************************************
//
//
// -- This is a parent command --
// Cypress.Commands.add('login', (email, password) => { ... })
//
//
// -- This is a child command --
// Cypress.Commands.add('drag', { prevSubject: 'element'}, (subject, options) => { ... })
//
//
// -- This is a dual command --
// Cypress.Commands.add('dismiss', { prevSubject: 'optional'}, (subject, options) => { ... })
//
//
// -- This will overwrite an existing command --
// Cypress.Commands.overwrite('visit', (originalFn, url, options) => { ... })

import firebase from "firebase/compat/app";
import "firebase/compat/auth";
import "firebase/compat/database";
import "firebase/compat/firestore";
import { attachCustomCommands } from "cypress-firebase";

const fbConfig = {
    apiKey: "AIzaSyC62LX6ufR_SiZJdLujUiXcR3Iy6xL_dqM",
    authDomain: "kora-2d67f.firebaseapp.com",
    projectId: "kora-2d67f",
    storageBucket: "kora-2d67f.appspot.com",
    messagingSenderId: "22175424879",
    appId: "1:22175424879:web:6e6e28e16852105350489a",
    // measurementId: "G-GRTPFV92RQ"  
};

firebase.initializeApp(fbConfig);

attachCustomCommands({ Cypress, cy, firebase });