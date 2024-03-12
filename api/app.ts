//npm packages
import express, { Application,Request, Response } from "express";
require("dotenv").config({
  path: "../.env",
});

// Custom Modules, Packages, Configs, etc.
import { connectDB } from "./databases/mongoDB";
import {redisClient} from './databases/redis';
import { getFixturesFromCache } from "./services/getFixturesFromCache";
import { getFixturesFromDatabase } from "./services/getFixturesFromDatabase";

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
    const page: number = parseInt(req.query.page as string) || 1;
    const limit: number = parseInt(req.query.limit as string) || 15

    const cacheData = await getFixturesFromCache(page, limit);
    if (cacheData) {
      return res.status(200).json({
        isCached: true,
        error: false,
        data: cacheData.fixtures,
        currentPage: page,
        totalPages: cacheData.totalPages
      });
    } else {
      const dbData = await getFixturesFromDatabase(page, limit);
      return res.status(200).json({
        isCached: false,
        error: false,
        data: dbData.fixtures,
        currentPage: page,
        totalPages: dbData.totalPages
      });
    }
  } catch (error) {
    console.error('Error fetching fixtures:', error);
    return res.status(500).json({error: true, message: 'Internal server error' });
  }
});


redisClient.on("connect", () => {
  console.log("Redis connected");
});
connectDB();
export default app;