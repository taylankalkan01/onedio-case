//npm packages
import express, { Application,Request, Response } from "express";
require("dotenv").config({
  path: "../.env",
});

// Custom Modules, Packages, Configs, etc.
import { connectDB } from "./database/mongoDB";
import Fixture from "./model/Fixture";

//Application
const app: Application = express();
app.set("trust proxy", 1);
app.disable("x-powered-by");
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

//healthcheck
app.get("/healthcheck", (_, res: Response) => {
  res.status(200).json({ error: false, message: "healthcheck" });
});

app.get("/fixtures", async (req: Request, res: Response) => {
    try {
        const limit = parseInt(req.query.limit as string) || 15;
        const page = parseInt(req.query.page as string) || 1;

        const startIndex = (page - 1) * limit;

        const fixtures = await Fixture.find().limit(limit).skip(startIndex).exec();
        const totalPages = Math.ceil(await Fixture.countDocuments() / limit);
        

        res.status(200).json({
            error:false,
            data: fixtures,
            currentPage: page,
            totalPages: totalPages
        });
    } catch (error) {
        console.error('Error fetching fixtures:', error);
        res.status(500).json({ error: 'Internal server error' });
    }
})



connectDB();
export default app;