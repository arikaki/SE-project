// Import the functions you need from the SDKs you need
import { initializeApp } from "firebase/app";
import { getAuth, GoogleAuthProvider } from "firebase/auth";
// import { getAnalytics } from "firebase/analytics";
// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
    apiKey: "AIzaSyC62LX6ufR_SiZJdLujUiXcR3Iy6xL_dqM",
    authDomain: "kora-2d67f.firebaseapp.com",
    projectId: "kora-2d67f",
    storageBucket: "kora-2d67f.appspot.com",
    messagingSenderId: "22175424879",
    appId: "1:22175424879:web:6e6e28e16852105350489a",
    measurementId: "G-GRTPFV92RQ"
  };

// Initialize Firebase
const app = initializeApp(firebaseConfig);
const auth = getAuth();
const provider = new GoogleAuthProvider();

export { auth, provider };