/**
 * Quiz API
 * Quiz App backend
 *
 * The version of the OpenAPI document: v1
 *
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

// May contain unused imports in some cases
// @ts-ignore
import { QuizQuestionAnswer } from "./quiz-question-answer";

/**
 *
 * @export
 * @interface QuizQuestion
 */
export interface QuizQuestion {
  /**
   * The sha1 of the whole quiz question
   * @type {string}
   * @memberof QuizQuestion
   */
  sha1?: string;
  /**
   * The question content
   * @type {string}
   * @memberof QuizQuestion
   */
  content?: string;
  /**
   *
   * @type {Array<QuizQuestionAnswer>}
   * @memberof QuizQuestion
   */
  questions?: Array<QuizQuestionAnswer>;
}