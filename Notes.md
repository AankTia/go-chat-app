# Go Programming Blueprints (Notes)

## Chapter 1: Chat Application with Web Sockets

### A simple web server 10

#### Separating views from logic using templates 12

##### Doing things once 14

##### Using your own handlers 14

#### Properly building and executing Go programs 15

### Modeling a chat room and clients on the server 15

#### Modeling the client 16

#### Modeling a room 19

#### Concurrency programming using idiomatic Go 19

#### Turning a room into an HTTP handler 20

#### Using helper functions to remove complexity 22

#### Creating and using rooms 23

### Building an HTML and JavaScript chat client 23

#### Getting more out of templates 25

### Tracing code to get a look under the hood 28

#### Writing a package using TDD 28

#### Interfaces 29

#### Unit tests 30

##### Red-green testing 32

#### Implementing the interface 34

##### Unexported types being returned to users 35

#### Using our new trace package

N/A

#### Making tracing optional

N/A

#### Clean package APIs 39

- In Go, adding documentation is as simple as adding comments to the line before each item.

### Summary

N/A

## Chapter 2: Adding User Accounts 41

### Handlers all the way down 42

...

### Making a pretty social sign-in page 45

- If you prefer to download and host your own copy of Bootstrap, you can do so. Keep the files in an `assets` folder and add the following call to your `main` function (it uses `http.Handle` to serve the assets via your application):

```go
http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("/path/to/assets/"))))
```

### Endpoints with dynamic paths

- Pattern matching for the `http` package in the Go standard library isn't the most comprehensive and fully featured implementation out there. For example, Ruby on Rails makes it much easier to have dynamic segments inside the path. You could map the route like this:

  `"auth/:action/:provider_name"`

  Rails then provides a data map (or dictionary) containing the values that it automatically extracted from the matched path. So if you visit `auth/login/google`, then `params[:provider_name]` would equal `google` and `params[:action]` would equal `login`.

- If you need to handle more advanced routing situations, you may want to consider using dedicated packages, such as `goweb`, `pat`, `routes`, or `mux`.

### Getting started with OAuth2

OAuth2 is an open authorization standard designed to allow resource owners to give clients delegated access to private data (such as wall posts or tweets) via an access token exchange handshake. Even if you do not wish to access the private data, OAuth2 is a great option that allows people to sign in using their existing credentials, without exposing those credentials to a third-party site. In this case, we are the third party, and we want to allow our users to sign in using services that support OAuth2.

From a user's point of view, the OAuth2 flow is as follows:

1. The user selects the provider with whom they wish to sign in to the client app.
2. The user is redirected to the provider's website (with a URL that includes the client app ID) where they are asked to give permission to the client app.
3. The user signs in from the OAuth2 service provider and accepts the permissions requested by the third-party application.
4. The user is redirected to the client app with a request code.
5. In the background, the client app sends the grant code to the provider, who sends back an authentication token.
6. The client app uses the access token to make authorized requests to the provider,
   such as to get user information or wall posts.

...

#### Open source OAuth2 packages

##### `gomniauth`

`gomniauth` (see https://github.com/stretchr/gomniauth). An open source Go alternative to Ruby's `omniauth` project, `gomniauth` provides a unified solution to access different OAuth2 services.

Some of the project dependencies of gomniauth are kept in Bazaar repositories, so you'll need to head over to http://wiki.bazaar.canonical.com to download them.

### Tell the authorization providers about your app

Before we ask an authorization provider to help our users sign in, we must tell them about our application. Most providers have some kind of web tool or console where you can create applications to kick this process off.

In order to identify the client application, we need to create a client ID and secret. Despite the fact that OAuth2 is an open standard, each provider has their own language and mechanism to set things up. Therefore, you will most likely have to play around with the user interface or the documentation to figure it out in each case.

If we host our application on a real domain, we have to create new client
IDs and secrets or update the appropriate URL fields on our authorization
providers to ensure that they point to the right place.

### Implementing external logging in

#### Logging in

#### Handling the response from the provider

Base64-encoding data ensures it won't contain any special or unpredictable characters, which is useful for situations such as passing data to a URL or storing it in a cookie. Remember that although Base64-encoded data looks encrypted, it is not you can easily decode Base64-encoded data back to the original text with little effort. There are online tools that do this for you.

### Presenting the user data

...

### Augmenting messages with additional data

...

### Summary

...

## Chapter 3: Three Ways to Implement Profile Pictures

...

### Avatars from the OAuth2 server

...

#### Getting the avatar URL

...

#### Transmitting the avatar URL

...

#### Adding the avatar to the user interface

...

#### Logging out

...

#### Making things prettier

...

### Implementing Gravatar

**Gravatar** is a web service that allows users to upload a single profile picture and associate it with their e-mail address in order to make it available from any website

#### Abstracting the avatar URL process 73

...

##### The auth service and the avatar's implementation

...

##### Using an implementation

...

##### The Gravatar implementation

...

### Uploading an avatar picture

...

#### User identification

...

#### An upload form

...

#### Handling the upload

...

#### Serving the images

...

#### The Avatar implementation for local files

...

##### Supporting different file types

...

#### Refactoring and optimizing our code

...

##### Replacing concrete types with interfaces

...

##### Changing interfaces in a test-driven way

...

##### Fixing the existing implementations 94

##### Global variables versus fields 95

##### Implementing our new design 96

##### Tidying up and testing 97

### Combining all three implementations 98

### Summary

## Chapter 4: Command-Line Tools to Find Domain Names 101

### Pipe design for command-line tools 102

### Five simple programs 102

#### Sprinkle 103

#### Domainify 107

#### Coolify

#### Synonyms 112

##### Using environment variables for configuration 113

##### Consuming a web API 113

##### Getting domain suggestions 117

#### Available 118

### Composing all five programs 122

#### One program to rule them all 123

### Summary

## Chapter 5: Building Distributed Systems and Working with Flexible Data 128

### The system design 129

#### The database design 130

### Installing the environment 131

#### Introducing NSQ 131

##### NSQ driver for Go 133

#### Introducing MongoDB 133

##### MongoDB driver for Go 134

#### Starting the environment 134

### Reading votes from Twitter 135

#### Authorization with Twitter 135

##### Extracting the connection 137

##### Reading environment variables 138

#### Reading from MongoDB 140

#### Reading from Twitter 142

##### Signal channels 144

#### Publishing to NSQ 146

#### Gracefully starting and stopping programs 148

#### Testing 150

### Counting votes 151

#### Connecting to the database 152

#### Consuming messages in NSQ 153

#### Keeping the database updated 155

#### Responding to Ctrl + C 157

### Running our solution 158

### Summary 159

## Chapter 6: Exposing Data and Functionality through a RESTful Data Web Service API 161

### RESTful API design 162

### Sharing data between handlers 163

#### Context keys 163

### Wrapping handler functions 165

#### API keys 165

#### Cross-origin resource sharing 166

### Injecting dependencies 167

### Responding 167

### Understanding the request 169

### Serving our API with one function 171

#### Using handler function wrappers 173

### Handling endpoints 173

#### Using tags to add metadata to structs 174

#### Many operations with a single handler 174

##### Reading polls 175

##### Creating a poll 178

##### Deleting a poll CORS support 179

#### Testing our API using curl 180

### A web client that consumes the API 182

#### Index page showing a list of polls 183

#### Creating a new poll 185

#### Showing the details of a poll 186

### Running the solution 189

### Summary 191

## Chapter 7: Random Recommendations Web Service 193

### The project overview 194

#### Project design specifics 195

### Representing data in code 197

#### Public views of Go structs 200

### Generating random recommendations 201

#### The Google Places API key 203

#### Enumerators in Go 203

##### Test-driven enumerator 205

#### Querying the Google Places API 209

#### Building recommendations 210

#### Handlers that use query parameters 212

#### CORS 213

#### Testing our API 214

##### Web application 216

### Summary 216

## Chapter 8: Filesystem Backup 218

### Solution design 219

#### The project structure 219

### The backup package 220

#### Considering obvious interfaces first 220

#### Testing interfaces by implementing them 221

#### Has the filesystem changed? 224

#### Checking for changes and initiating a backup 226

##### Hardcoding is OK for a short while 228

### The user command-line tool 229

#### Persisting small data 230

#### Parsing arguments 231

##### Listing the paths 232

###### String representations for your own types 232

##### Adding paths 233

##### Removing paths 233

#### Using our new tool 234

### The daemon backup tool 235

#### Duplicated structures 237

#### Caching data 237

#### Infinite loops 238

#### Updating filedb records 239

### Testing our solution 240

### Summary 242

## Chapter 9: Building a Q&A Application for Google App Engine 243

### The Google App Engine SDK for Go 244

#### Creating your application 245

#### App Engine applications are Go packages 246

#### The app.yaml file 246

#### Running simple applications locally 247

#### Deploying simple applications to Google App Engine 249

#### Modules in Google App Engine 250

##### Specifying modules 251

##### Routing to modules with dispatch.yaml 252

### Google Cloud Datastore 252

#### Denormalizing data 253

### Entities and data access 255

#### Keys in Google Cloud Datastore 256

#### Putting data into Google Cloud Datastore 257

#### Reading data from Google Cloud Datastore 259

### Google App Engine users 259

#### Embedding denormalized data 261

### Transactions in Google Cloud Datastore 262

#### Using transactions to maintain counters 263

#### Avoiding early abstraction 267

### Querying in Google Cloud Datastore 267

### Votes 269

#### Indexing 270

#### Embedding a different view of entities 271

### Casting a vote 273

#### Accessing parents via datastore.Key 274

#### Line of sight in code 274

### Exposing data operations over HTTP 277

#### Optional features with type assertions 277

#### Response helpers 278

#### Parsing path parameters 279

#### Exposing functionality via an HTTP API 281

##### HTTP routing in Go 281

#### Context in Google App Engine 282

#### Decoding key strings 283

##### Using query parameters 285

##### Anonymous structs for request data 286

##### Writing self-similar code 287

##### Validation methods that return an error 288

#### Mapping the router handlers 289

### Running apps with multiple modules 290

#### Testing locally 290

##### Using the admin console 291

###### Automatically generated indexes 292

### Deploying apps with multiple modules 292

### Summary 293

## Chapter 10: Micro-services in Go with the Go kit Framework 294

### Introducing gRPC 296

### Protocol buffers 297

#### Installing protocol buffers 298

#### Protocol buffers language 298

#### Generating Go code 300

### Building the service 301

#### Starting with tests 302

#### Constructors in Go 303

#### Hashing and validating passwords with bcrypt 304

### Modeling method calls with requests and responses 305

#### Endpoints in Go kit 307

##### Making endpoints for service methods 308

##### Different levels of error 309

##### Wrapping endpoints into a Service implementation 309

### An HTTP server in Go kit 311

### A gRPC server in Go kit 312

#### Translating from protocol buffer types to our types 313

### Creating a server command 315

#### Using Go kit endpoints 318

#### Running the HTTP server 318

#### Running the gRPC server 319

#### Preventing a main function from terminating immediately 320

#### Consuming the service over HTTP 320

### Building a gRPC client 321

#### A command-line tool to consume the service 323

#### Parsing arguments in CLIs 324

#### Maintaining good line of sight by extracting case bodies 325

#### Installing tools from the Go source code 326

### Rate limiting with service middleware 327

#### Middleware in Go kit 328

#### Manually testing the rate limiter 330

#### Graceful rate limiting 331

### Summary 332

## Chapter 11: Deploying Go Applications Using Docker 333

### Using Docker locally 334

#### Installing Docker tools 334

#### Dockerfile 334

#### Building Go binaries for different architectures 335

#### Building a Docker image 336

#### Running a Docker image locally 337

#### Inspecting Docker processes 338

#### Stopping a Docker instance 339

### Deploying Docker images 339

#### Deploying to Docker Hub 339

### Deploying to Digital Ocean 341

#### Creating a droplet 341

#### Accessing the droplet's console 344

#### Pulling Docker images 346

#### Running Docker images in the cloud

#### Accessing Docker images in the cloud

### Summary 349

## Appendix: Good Practices for a Stable Go Environment 348

### Installing Go 350

### Configuring Go 351

#### Getting GOPATH right 352

### Go tools 353

### Cleaning up, building, and running tests on save 355

### Integrated developer environments 356

#### Sublime Text 3 356

### Visual Studio Code 359

### Summary 362

## Index 363
