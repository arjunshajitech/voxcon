package constant

const Charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var JwtSecretKey = []byte("my_secret_key")

var PlayerLeftWscl = "player_left_ws_connection_lost"
var PlayerLeftSpd = "player_left_sender_peer_disconnect"
var PlayerLeftRpd = "player_left_receiver_peer_disconnect"

var OfferSDP = "offer_sdp"
var AnswerSDP = "answer_sdp"

var DefaultGameID = "Game-0"
