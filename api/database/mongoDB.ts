import mongoose, { ConnectOptions } from "mongoose";

const MONGO_URL = process.env.MONGO_URL || "";
const dbName = 'football';
mongoose.set("strictQuery", false);

export const connectDB = async () => {
  await mongoose
    .connect(`${MONGO_URL}${dbName}`, {
      dbName: dbName,
    } as ConnectOptions)
    .then((res) => {
      console.log("Database Connected Successfuly.", res.connection.host);
    })
    .catch((err) => {
      console.log("Database connection error: ", err);
    });
};