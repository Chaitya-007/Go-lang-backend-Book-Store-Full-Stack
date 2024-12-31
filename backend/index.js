import express from "express";
// import { PORT, mongoDBURL } from "./config.js";
import mongoose from "mongoose";
import { Book } from "./models/bookModel.js";
import booksRoute from "./routes/booksRoute.js";
import cors from "cors";
import dotenv from "dotenv";
import path from "path";
dotenv.config({});

const PORT = process.env.PORT || 5000;
const mongoDBURL = process.env.MONGODB_URL;

// Path of directory is fetched
const _dirname = path.resolve();

const app = express();

// Middleware for parsing request body
app.use(express.json());

// Middleware for handling CORS policy
// Option 1: Allow all Origins with Default of cors(*)
app.use(cors());
// Option 2: Allow Custom Origin
// app.use(
//   cors({
//     origin: "http://localhost:3000",
//     methods: ["GET", "POST", "PUT", "DELETE"],
//     allowedHeaders: ["Content-Type"],
//   })
// );

// app.get("/", (request, response) => {
//   console.log(request);
//   return response.status(234).send("Welcome to MERN Stack Tutorial");
// });

// This will add /books to all the routes in booksRoute in request as prefix
app.use("/books", booksRoute);

app.use(express.static(path.join(_dirname, "/frontend/dist")));
app.get("*", (req, res) => {
  res.sendFile(path.resolve(_dirname, "/frontend/dist/index.html"));
});

mongoose
  .connect(mongoDBURL)
  .then(() => {
    console.log("Connected to MongoDB");
    app.listen(PORT, () => {
      console.log(`App is listening to port : ${PORT}`);
    });
  })
  .catch((error) => {
    console.log(error);
  });
