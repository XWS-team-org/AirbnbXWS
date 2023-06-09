using Neo4jClient;
using ReviewService.Handlers;
using ReviewService.Repository;
using ReviewService.Service;

var builder = WebApplication.CreateBuilder(args);

// Additional configuration is required to successfully run gRPC on macOS.
// For instructions on how to configure Kestrel and gRPC clients on macOS, visit https://go.microsoft.com/fwlink/?linkid=2099682

// Add services to the container.
builder.Services.AddGrpc();
builder.Services.AddScoped(typeof(IRepository<>), typeof(Neo4jRepository<>));
builder.Services.AddScoped<IReviewService, RatingService>();
var app = builder.Build();


app.UseRouting();
// Configure the HTTP request pipeline.
app.UseEndpoints(endpoints =>
{
    endpoints.MapGrpcService<ReviewHandler>();
}
);
app.MapGet("/", () => "Communication with gRPC endpoints must be made through a gRPC client. To learn how to create a client, visit: https://go.microsoft.com/fwlink/?linkid=2086909");

app.Run();
