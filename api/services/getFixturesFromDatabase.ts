import Fixture from "../model/Fixture";

export async function getFixturesFromDatabase(page: number, limit: number): Promise<any> {
    try {
      const startIndex: number = (page - 1) * limit;
      const fixtures = await Fixture.find().limit(limit).skip(startIndex).exec();
      const totalCount = await Fixture.countDocuments();
      const totalPages = Math.ceil(totalCount / limit);
      return { fixtures, totalPages };
    } catch (error) {
      throw error;
    }
  }
  