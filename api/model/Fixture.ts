import mongoose, { Document } from "mongoose";


interface Fixture extends Document {
    Date: string;
    HomeTeam: string;
    AwayTeam: string;
    FTHG: number;
    FTAG: number;
    Referee: string;
}

const fixtureSchema = new mongoose.Schema<Fixture>(
    {
        Date: {
            type: String,
            required: true,
        },
        HomeTeam: {
            type: String,
            required: true,
        },
        AwayTeam: {
            type: String,
            required: true,
        },
        FTHG: {
            type: Number,
            required: true,
        },
        FTAG: {
            type: Number,
            required: true,
        },
        Referee: {
            type: String,
            required: true,
        },
    },
  {
    versionKey: false,
    timestamps: true,
    collection: 'fixtures'
    
  },
);

const Fixture = mongoose.model<Fixture>("fixtures", fixtureSchema);

export default Fixture;