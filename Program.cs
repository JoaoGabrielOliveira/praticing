var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
builder.Services.AddControllersWithViews();

var app = builder.Build();

// Configure the HTTP request pipeline.
if (!app.Environment.IsDevelopment())
{
    app.UseExceptionHandler("/Home/Error");
    // The default HSTS value is 30 days. You may want to change this for production scenarios, see https://aka.ms/aspnetcore-hsts.
    app.UseHsts();
    Environment.SetEnvironmentVariable("connectionString", "Server=localhost:1433;User Id=;Password=Pass1234;");
}

app.UseHttpsRedirection();

app.UseRouting();

app.UseAuthorization();

app.UseEndpoints(endpoints =>
{
    endpoints.MapControllerRoute(
           name: "default",
           pattern: "/student",
           defaults: new {controller="Student", action="Index" });

    endpoints.MapControllerRoute(
           name: "default",
           pattern: "/student/{id}",
           defaults: new {controller="Student", action="Get"});
    
});

app.Run();
