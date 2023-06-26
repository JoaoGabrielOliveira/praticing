namespace Database;

using Microsoft.EntityFrameworkCore;
using Models;

public class BaseContext<T, ID> : DbContext where T : Model<ID> {
    public DbSet<T> Data {get; set;} = null!;

    protected override void OnConfiguring(DbContextOptionsBuilder options){
        options.UseSqlServer(Environment.GetEnvironmentVariable("connectionString"));
    }
}